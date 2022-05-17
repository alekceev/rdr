package starter

import (
	"context"
	"sync"
)

type App struct {
}

func NewApp() *App {
	a := &App{}
	return a
}

type HTTPServer interface {
	Start()
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start()
	<-ctx.Done()
	hs.Stop()
}
