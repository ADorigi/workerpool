package workerpool

import (
	"context"
	"log"
	"sync"

	"github.com/google/uuid"
)

type ResultWorker struct {
	id      uuid.UUID
	results <-chan Result
	wg      *sync.WaitGroup
}

func NewResultWorker(
	results <-chan Result,
	wg *sync.WaitGroup,
) ResultWorker {
	return ResultWorker{
		id:      uuid.New(),
		results: results,
		wg:      wg,
	}
}

func (rw *ResultWorker) Start(ctx context.Context) {
	go func() {
		for result := range rw.results {
			log.Printf("Task: %s :: err: %s", result.GetID(), result.GetErr())
			rw.wg.Done()
		}
	}()
}
