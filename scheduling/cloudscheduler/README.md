# Google Cloud Scheduler

## Infrastructure

Creates a job that makes a POST call to an HTTP endpoint with a `generic.RawObject` body at the scheduled times.

Environment variables:

- `GOOGLE_APPLICATION_CREDENTIALS_FILENAME`: path to the a Google Cloud service account JSON credentials file
- `GOOGLE_CLOUD_SCHEDULER_PROJECT_ID`: ID of the project where the job is being created
- `GOOGLE_CLOUD_SCHEDULER_LOCATION_ID`: ID of the location where the job should be created (e.g. southamerica-east1)
- `GOOGLE_CLOUD_SCHEDULER_HTTP_TARGET_URL`: URL to the HTTP target of the scheduler jobs
- `GOOGLE_CLOUD_SCHEDULER_JOB_SCHEDULE`: cron tab representing the schedule for the job (e.g. "* * * * *" means "every minute")
- `GOOGLE_CLOUD_SCHEDULER_JOB_TIMEZONE`: must be a time zone name from the tz database (e.g. America/Sao_Paulo)

## Server

Creates a `generic.HTTPServer` to receive the calls made by a Google Cloud Scheduler Job.
