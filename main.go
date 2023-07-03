package main

import (
	gpntron "go-gpn-tron/internal/gpn-tron"
	"go-gpn-tron/internal/strategies"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// const host = "gpn-tron.duckdns.org:4000"
const host = "localhost:4000"

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	client := gpntron.NewClient(gpntron.ClientOptions{
		Host:            host,
		Username:        "1trick",
		Password:        "1337424242421337hahafunnynumbers",
		TimeoutDuration: 5 * time.Minute,
	})

	go func() {
		<-sigs
		client.Disconnect()
	}()

	client.Run(&strategies.Random{})
}
