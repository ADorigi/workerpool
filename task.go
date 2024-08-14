package workerpool

import (
	"context"

	"github.com/google/uuid"
)

type TaskProperties struct {
	ID          uuid.UUID
	Description string
}

type Task interface {
	Properties() TaskProperties
	Run(ctx context.Context) error
}
