package handler

import (
	"context"
	"net/url"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	r := &Handlers{}
	return r
}

func (rt *Handlers) Admin(ctx context.Context, url url.URL) error {
	// TODO
	return nil
}
