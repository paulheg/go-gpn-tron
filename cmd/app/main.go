package main

import (
	"context"
	gpntron "go-gpn-tron/internal/gpn-tron"
	"go-gpn-tron/internal/strategies"
	"go-gpn-tron/internal/visualize"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const host = "gpn-tron.duckdns.org:4000"

// const host = "151.216.74.213:4000"

// const host = "localhost:4000"

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	client := gpntron.NewClient(gpntron.ClientOptions{
		Host:            host,
		Username:        "1trick",
		Password:        "1337424242421337hahafunnynumbers",
		TimeoutDuration: 5 * time.Minute,
	})

	go func() {
		<-sigs
		cancel()
	}()

	strat := strategies.Random{}

	if err := client.Run(ctx, visualize.NewConsoleVisualizer(&strat, &strat, &strat.Arena)); err != nil {
		log.Fatalf("error running client: %s", err)
	}
}
