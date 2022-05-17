package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"rdr/api/handler"
	"rdr/api/router"
	"rdr/api/server"
	"rdr/app/config"
	"rdr/app/starter"
)

func main() {
	// Config
	conf, err := config.Get()
	if err != nil {
		log.Fatalf("Error parsing config: %v\n", err)
	}

	// Set tz
	if tz := os.Getenv("TZ"); tz != "" {
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("error loading location '%s': %v\n", tz, err)
		}
	}

	// output current time zone
	tnow := time.Now()
	tz, _ := tnow.Zone()
	log.Printf("Local time zone %s. Service started at %s", tz,
		tnow.Format("2006-01-02T15:04:05.000 MST"))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	a := starter.NewApp()
	h := handler.NewHandlers()

	rh := router.NewRouter(h)
	srv := server.NewServer(conf, rh)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go a.Serve(ctx, wg, srv)

	<-ctx.Done()
	cancel()
	wg.Wait()
}
