package client

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"poc/shared/generic"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateTasks(credentialsFilename, queueName, workerURL string, maxConcurrentRequests, tasksToBeCreated int) {
	ctx := context.Background()

	taskCreator, err := newTaskCreator(ctx, credentialsFilename, queueName, workerURL)
	if err != nil {
		log.Fatalf("failed to create Cloud Tasks client: %v\n", err)
	}

	now := time.Now()

	semaphore := make(chan struct{}, maxConcurrentRequests)

	var g errgroup.Group

	for i := 0; i < tasksToBeCreated; i++ {
		semaphore <- struct{}{}

		g.Go(func() error {
			defer func() {
				<-semaphore
			}()

			taskID, err := taskCreator.createTask(ctx, generic.Object)
			if err != nil {
				log.Fatalf("failed to create task: %v\n", err)
			}

			log.Printf("task %s created successfully\n", taskID)
			return nil
		})
	}

	g.Wait()

	took := time.Since(now)

	log.Printf("created %d tasks in %s\n", tasksToBeCreated, took)
}

type taskCreator struct {
	client    *cloudtasks.Client
	queueName string
	workerURL string
}

func newTaskCreator(ctx context.Context, credentialsFilename, queueName, workerURL string) (*taskCreator, error) {
	cctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := cloudtasks.NewClient(cctx, option.WithCredentialsFile(credentialsFilename))
	if err != nil {
		return nil, err
	}

	return &taskCreator{
		client:    client,
		queueName: queueName,
		workerURL: workerURL,
	}, nil
}

func (c *taskCreator) createTask(ctx context.Context, data generic.Struct) (string, error) {
	cctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	request := &cloudtaskspb.CreateTaskRequest{
		Parent: c.queueName,
		Task: &cloudtaskspb.Task{
			MessageType: &cloudtaskspb.Task_HttpRequest{
				HttpRequest: &cloudtaskspb.HttpRequest{
					Url:        c.workerURL,
					HttpMethod: cloudtaskspb.HttpMethod_POST,
					Headers: map[string]string{
						"Content-Type": "application/json",
					},
					Body: body,
				},
			},
			CreateTime: timestamppb.Now(),
		},
	}

	task, err := c.client.CreateTask(cctx, request)
	if err != nil {
		return "", err
	}

	return task.GetName(), nil
}
