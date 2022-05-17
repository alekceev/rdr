package handler

import (
	"context"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	r := &Handlers{}
	return r
}

func (rt *Handlers) Admin(ctx context.Context) error {
	// TODO
	return nil
}
