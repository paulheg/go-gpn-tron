package main

import (
	"context"
	gpntron "go-gpn-tron/internal/gpn-tron"
	"go-gpn-tron/internal/strategies"
	"go-gpn-tron/internal/visualize"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const host = "gpn-tron.duckdns.org:4000"

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := gpntron.NewClient(gpntron.ClientOptions{
		Host:            host,
		Username:        "1trick",
		Password:        "1337424242421337hahafunnynumbers",
		TimeoutDuration: 5 * time.Minute,
	})

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("go-gpn-tron")

	game := visualize.NewEbitenVisualizer(&strategies.Random{})

	go func() {
		<-sigs
		log.Print("Closing everything.")
		cancel()
		game.Close()
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		client.Run(ctx, game)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := ebiten.RunGame(game)
		log.Printf("Window Cancelled: %v", err)
		cancel()
	}()

	wg.Wait()
	log.Println("exited")
}
