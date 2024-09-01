# Job Scheduler Project

This project is a simple job scheduler that allows you to execute jobs concurrently with a specified maximum fanout and retry limit. The scheduler is designed to handle job failures and retries them up to the specified limit.

## Features

- Concurrent job execution with a maximum fanout
- Job retry mechanism with a specified retry limit
- Shutdown functionality to stop the scheduler

## Usage

To use the job scheduler, follow these steps:

1. Create a new instance of the `JobScheduler` with the desired maximum fanout and retry limit.
2. Start the scheduler by calling the `Run` method.
3. Add jobs to the scheduler using the `AddJob` method.
4. Shutdown the scheduler when all jobs have been added using the `Shutdown` method.

## Example

The `main.go` file in this project includes an example of how to use the job scheduler. It creates a scheduler with a maximum fanout of 3 and a retry limit of 3, adds 10 simulated jobs, and then shuts down the scheduler.

## Simulated Job

The `simulatedJob` function is a placeholder for actual job logic. It simulates a job that randomly fails with a probability of 30%. This allows for demonstration of the retry mechanism.

## License

This project is licensed under the MIT License.
