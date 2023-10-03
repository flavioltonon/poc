package main

import (
	"context"
	"log"
	"os"
	"path"
	"poc/shared/generic"
	"time"

	scheduler "cloud.google.com/go/scheduler/apiv1"
	"cloud.google.com/go/scheduler/apiv1/schedulerpb"
	"google.golang.org/api/option"
)

var (
	// credentialsFilename is the path to the a Google Cloud service account JSON credentials file
	credentialsFilename = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_FILENAME")

	// projectID is the ID of the project where the job is being created
	projectID = os.Getenv("GOOGLE_CLOUD_SCHEDULER_PROJECT_ID")

	// locationID is the ID of the location where the job should be created (e.g. southamerica-east1)
	locationID = os.Getenv("GOOGLE_CLOUD_SCHEDULER_LOCATION_ID")

	// httpTargetURL is the URL to the HTTP target of the scheduler jobs
	httpTargetURL = os.Getenv("GOOGLE_CLOUD_SCHEDULER_HTTP_TARGET_URL")

	// jobSchedule is the cron tab representing the schedule for the job (e.g. "* * * * *" means "every minute")
	jobSchedule = os.Getenv("GOOGLE_CLOUD_SCHEDULER_JOB_SCHEDULE")

	// jobTimezone must be a time zone name from the tz database (e.g. America/Sao_Paulo)
	jobTimezone = os.Getenv("GOOGLE_CLOUD_SCHEDULER_JOB_TIMEZONE")
)

func main() {
	ctx := context.Background()

	provider, err := newProvider(ctx, credentialsFilename, projectID, locationID)
	if err != nil {
		log.Fatalf("failed to create provider: %v\n", err)
	}

	defer provider.stop()

	jobID, err := provider.createJob(ctx, &schedulerpb.Job{
		Target: &schedulerpb.Job_HttpTarget{
			HttpTarget: &schedulerpb.HttpTarget{
				Uri:        httpTargetURL,
				HttpMethod: schedulerpb.HttpMethod_POST,
				Body:       generic.RawObject,
			},
		},
		Schedule: jobSchedule,
		TimeZone: jobTimezone,
	})
	if err != nil {
		log.Fatalf("failed to create job: %v\n", err)
	}

	log.Printf("job %s created\n", jobID)
}

type provider struct {
	client     *scheduler.CloudSchedulerClient
	projectID  string
	locationID string
}

func newProvider(ctx context.Context, credentialsFilename, projectID, locationID string) (*provider, error) {
	cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	client, err := scheduler.NewCloudSchedulerClient(cctx, option.WithCredentialsFile(credentialsFilename))
	if err != nil {
		return nil, err
	}

	return &provider{
		client:     client,
		projectID:  projectID,
		locationID: locationID,
	}, nil
}

func (p *provider) createJob(ctx context.Context, jobTemplate *schedulerpb.Job) (string, error) {
	cctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	request := &schedulerpb.CreateJobRequest{
		Parent: path.Join("projects", p.projectID, "locations", p.locationID),
		Job:    jobTemplate,
	}

	job, err := p.client.CreateJob(cctx, request)
	if err != nil {
		return "", err
	}

	return job.Name, nil
}

func (p *provider) stop() error {
	return p.client.Close()
}
