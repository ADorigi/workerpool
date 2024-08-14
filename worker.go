package workerpool

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
)

type Worker struct {
	id        uuid.UUID
	taskQueue <-chan Task
	results   chan<- Result
	wg        *sync.WaitGroup
}

func NewWorker(
	taskQueue <-chan Task,
	results chan<- Result,
	wg *sync.WaitGroup,
) Worker {
	return Worker{
		id:        uuid.New(),
		taskQueue: taskQueue,
		results:   results,
		wg:        wg,
	}
}

func (w *Worker) Start(ctx context.Context) {
	go func() {
		for task := range w.taskQueue {
			log.Printf("Running task: %s", task.Properties().Description)
			err := task.Run(ctx)
			w.results <- *NewResult(
				task.Properties().ID,
				err,
			)
			w.wg.Done()
		}
	}()
}
