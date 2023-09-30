package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path"
	"poc/google/cloudtasks/shared"
	"time"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	var (
		credentialsFilename   = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_FILENAME")
		queueName             = os.Getenv("GOOGLE_CLOUD_TASKS_QUEUE_NAME")
		maxConcurrentRequests = 100
		tasksToBeCreated      = 10000
		workerURL             = os.Getenv("GOOGLE_CLOUD_TASKS_WORKER_URL")
	)

	ctx := context.Background()

	client, err := createCloudTasksClient(ctx, credentialsFilename, queueName, workerURL)
	if err != nil {
		log.Fatalf("failed to create Cloud Tasks client: %v\n", err)
	}

	defer client.close()

	now := time.Now()

	semaphore := make(chan struct{}, maxConcurrentRequests)

	var g errgroup.Group

	for i := 0; i < tasksToBeCreated; i++ {
		semaphore <- struct{}{}

		g.Go(func() error {
			defer func() {
				<-semaphore
			}()

			data := shared.TaskData{
				BatchID: uuid.NewString(),
				EntryID: uuid.NewString(),
			}

			taskID, err := client.createTask(ctx, data)
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

type gRPCClient struct {
	client        *cloudtasks.Client
	queueName     string
	taskIDFactory taskIDFactory
	workerURL     string
}

func createCloudTasksClient(ctx context.Context, credentialsFilename, queueName, workerURL string) (*gRPCClient, error) {
	cctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := cloudtasks.NewClient(cctx, option.WithCredentialsFile(credentialsFilename))
	if err != nil {
		return nil, err
	}

	return &gRPCClient{
		client:        client,
		queueName:     queueName,
		taskIDFactory: taskIDFactoryFunc(uuid.NewString),
		workerURL:     workerURL,
	}, nil
}

func (c *gRPCClient) close() error {
	return c.client.Close()
}

func (c *gRPCClient) createTask(ctx context.Context, data shared.TaskData) (string, error) {
	cctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	taskID := c.taskIDFactory.createTaskID()

	body, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	request := &cloudtaskspb.CreateTaskRequest{
		Parent: c.queueName,
		Task: &cloudtaskspb.Task{
			Name: c.createTaskName(taskID),
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

	if _, err := c.client.CreateTask(cctx, request); err != nil {
		return "", err
	}

	return taskID, nil
}

func (c *gRPCClient) createTaskName(taskID string) string {
	return path.Join(c.queueName, "tasks", taskID)
}

type taskIDFactory interface {
	createTaskID() string
}

type taskIDFactoryFunc func() string

func (fn taskIDFactoryFunc) createTaskID() string { return fn() }
