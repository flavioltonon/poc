# Google Cloud Tasks

## Client

Creates tasks on a Google Cloud Tasks queue.

Inputs:

- credentialsFilename: path to a Google Cloud service account JSON credentials file
- queueName: name of the queue the tasks should be created on. Pattern `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID`
- maxConcurrentRequests: number of requests that are allowed to be made concurrently to Google Cloud Tasks
- tasksToBeCreated: number of tasks expected to be created
- workerURL: URL of the worker that should receive the task

## Server

Exposes an endpoint for receiving tasks from a Google Cloud Tasks queue and calculates the throughput of tasks received (RPS).

Inputs: None

P.S: Google Cloud Tasks requires a public endpoint to send tasks. One way of achieving this locally is to expose your server with [ngrok](https://ngrok.com/).