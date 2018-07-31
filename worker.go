package main

import (
	"context"
	"log"

	"github.com/google/go-cloud/blob"
	"github.com/pkg/errors"
)

var (
	MaxWorkers = 2 //os.Getenv("MAX_WORKERS")
	MaxQueue   = 2 //os.Getenv("MAX_QUEUE")
)

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
	bucket     *blob.Bucket
	ctx        context.Context
}

func NewWorker(workerPool chan chan Job) Worker {
	ctx := context.Background()
	//Setup AWS using go cloud
	//One session per worker
	bucket, err := SetupAws(ctx, S3StorageBucketName)
	if err != nil {
		panic(errors.Wrap(err, "Unable to setup AWS Sessions"))
	}
	return Worker{
		WorkerPool: workerPool, // Holding to worker pool recived from Dispatcher, which is channel
		//of Job Queue ( channel of jobs )
		JobChannel: make(chan Job), // Only one job at a time to process by worker
		quit:       make(chan bool),
		bucket:     bucket,
		ctx:        ctx,
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for { // Run forever
			// register the current worker job queue to dispatcher job queue
			//REM: This is a for ever loop of worker and once a job is placed on to workers queue
			// below code( select block ) will execute and once completed, here the worker is placing his queue back to
			// worker pool to get another job
			w.WorkerPool <- w.JobChannel

			select { // Synchronously wait

			case job := <-w.JobChannel: // Once a job is placed onto this workers job queue pick it up and start working
				// This works in conjunction with worker pool above. i.e. we sharing the workers queue to the worker pool.
				if err := job.Payload.Upload(w.bucket, w.ctx); err != nil {
					log.Printf("Error uploading : %s", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
