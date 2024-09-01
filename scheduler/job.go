package scheduler

import (
	"log"
	"sync"
	"time"
)

// Job represents a job that can be scheduled and executed.
type Job func() error

// JobScheduler manages job scheduling with concurrency control.
type JobScheduler struct {
	jobQueue   chan Job
	maxFanout  int
	wg         sync.WaitGroup
	retryLimit int
}

// NewJobScheduler creates a new JobScheduler with a specified fanout limit.
func NewJobScheduler(maxFanout int, retryLimit int) *JobScheduler {
	return &JobScheduler{
		jobQueue:   make(chan Job),
		maxFanout:  maxFanout,
		retryLimit: retryLimit,
	}
}

// AddJob adds a job to the scheduler's queue.
func (s *JobScheduler) AddJob(job Job) {
	s.jobQueue <- job
}

// Run starts the job scheduler, processing jobs concurrently.
func (s *JobScheduler) Run() {
	for i := 0; i < s.maxFanout; i++ {
		s.wg.Add(1)
		go s.worker()
	}
}

// worker processes jobs from the job queue.
func (s *JobScheduler) worker() {
	defer s.wg.Done()
	for job := range s.jobQueue {
		s.executeJobWithRetry(job, s.retryLimit)
	}
}

// executeJobWithRetry runs a job and retries if it fails.
func (s *JobScheduler) executeJobWithRetry(job Job, retries int) {
	for retries > 0 {
		err := job()
		if err == nil {
			log.Printf("Job completed successfully")
			return
		}
		log.Printf("Job failed with error: %v. Retrying... (%d retries left)\n", err, retries-1)
		retries--
		time.Sleep(1 * time.Second) // Simulate delay before retry
	}
	log.Printf("Job failed after maximum retries")
}

// Shutdown stops the job scheduler and waits for all workers to finish.
func (s *JobScheduler) Shutdown() {
	close(s.jobQueue)
	s.wg.Wait()
	log.Println("All jobs completed, scheduler shutting down.")
}
