package main

import (
	"fmt"
	scheduler "jobscheduler/scheduler"
	"math/rand"
	"time"
)

// Simulated job function that randomly fails
func simulatedJob() error {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if rand.Float32() > 0.7 {
		return fmt.Errorf("simulated job failure")
	}
	return nil
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().Unix())

	// Create a new JobScheduler with a max fanout of 3 and a retry limit of 3
	jobScheduler := scheduler.NewJobScheduler(3, 3)

	// Start the scheduler
	go jobScheduler.Run()

	// Add jobs to the scheduler
	for i := 0; i < 10; i++ {
		jobScheduler.AddJob(simulatedJob)
	}

	// Shutdown the scheduler after all jobs have been added
	jobScheduler.Shutdown()
}
