package main

import (
	"log"
	"os"
	"poc/google/cloudtasks/client"
	"poc/google/cloudtasks/server"
)

var (
	credentialsFilename   = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_FILENAME")
	serviceAccountEmail   = os.Getenv("GOOGLE_CLOUD_SERVICE_ACCOUNT_EMAIL")
	maxConcurrentRequests = 100
	queueName             = os.Getenv("GOOGLE_CLOUD_TASKS_QUEUE_NAME")
	showBodyParsingLogs   = false
	tasksToBeCreated      = 1000
	workerURL             = os.Getenv("GOOGLE_CLOUD_TASKS_WORKER_URL")
)

func main() {
	client.CreateTasks(credentialsFilename, serviceAccountEmail, queueName, workerURL, maxConcurrentRequests, tasksToBeCreated)

	if err := server.ListenAndServe(showBodyParsingLogs); err != nil {
		log.Fatalf("failed to listen and serve: %v\n", err)
	}
}
