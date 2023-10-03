package main

import (
	"os"
	"poc/google/cloudtasks/client"
	"poc/google/cloudtasks/server"
)

var (
	credentialsFilename   = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_FILENAME")
	maxConcurrentRequests = 100
	queueName             = os.Getenv("GOOGLE_CLOUD_TASKS_QUEUE_NAME")
	showBodyParsingLogs   = false
	tasksToBeCreated      = 1000
	workerURL             = os.Getenv("GOOGLE_CLOUD_TASKS_WORKER_URL")
)

func main() {
	client.CreateTasks(credentialsFilename, queueName, workerURL, maxConcurrentRequests, tasksToBeCreated)
	server.ListenAndServe(showBodyParsingLogs)
}
