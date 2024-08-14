package workerpool

import "github.com/google/uuid"

type Result struct {
	taskID  uuid.UUID
	success bool
	err     error
}

func NewResult(taskID uuid.UUID, err error) *Result {
	return &Result{
		taskID:  taskID,
		success: false,
		err:     err,
	}
}

func (r *Result) GetID() string {
	return r.taskID.String()
}

func (r *Result) GetErr() string {
	if r.err != nil {
		return r.err.Error()
	}
	return ""
}
