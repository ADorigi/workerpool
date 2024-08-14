package workerpool

import (
	"context"
	"sync"
)

type WorkerPool struct {
	taskQueue     chan Task
	resultChan    chan Result
	maxConcurrent int
	wg            *sync.WaitGroup
}

func NewWorkerPool(maxConcurrent int) *WorkerPool {
	return &WorkerPool{
		taskQueue:     make(chan Task, 10),
		resultChan:    make(chan Result, 10),
		maxConcurrent: maxConcurrent,
		wg:            &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.maxConcurrent; i++ {
		worker := NewWorker(wp.taskQueue, wp.resultChan, wp.wg)
		worker.Start(ctx)
	}
	resultWorker := NewResultWorker(wp.resultChan, wp.wg)
	resultWorker.Start(ctx)
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.wg.Add(2) // task and result
	wp.taskQueue <- task
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}
