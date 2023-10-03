# Google Cloud Scheduler

## Infrastructure

Use `google/cloudscheduler/infrastructure/up` to create a job that makes a POST call to an HTTP endpoint with a `generic.RawObject` body at the scheduled times.

## Server

Run `go run google/cloudscheduler/server/main.go` to create a `generic.HTTPServer` to receive the calls made by a Google Cloud Scheduler Job.
