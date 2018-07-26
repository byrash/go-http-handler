package main

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerPool chan chan Job
	maxWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	// Creating worker pool channel
	// This channel for channel of jobs ( Job queue ) will be passed to each worker
	// Each worker registers his job queue tito this channel
	workerPool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: workerPool, maxWorkers: maxWorkers}
}

func (d *Dispatcher) Run(jobQueue chan Job) {
	// starting n number of workers
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch(jobQueue)
}

func (d *Dispatcher) dispatch(jobQueue chan Job) {
	for {
		select {
		case job := <-jobQueue:
			// a job request has been received; Fire off and wait for worker
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}
