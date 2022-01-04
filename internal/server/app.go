package server

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"weltraumschaf.de/battleship/internal/server/router"
)

const httpPort = 10000

func Create() *cli.App {
	return &cli.App{
		Name:    "battleship",
		Version: "1.0.0",
		Authors: []*cli.Author{
			{
				Name:  "Sven Strittmatter",
				Email: "ich@weltraumschaf.de",
			},
		},
		Action: Execute,
	}
}

func Execute(c *cli.Context) error {
	log.Println("Starting battleship server...")
	baseCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	address := fmt.Sprintf("%s:%d", "localhost", httpPort)
	httpServer := &http.Server{
		Addr:        address,
		Handler:     router.NewConfiguredRouter(),
		BaseContext: func(_ net.Listener) context.Context { return baseCtx },
	}

	// Run server
	go func() {
		log.Println("Running HTTP server.")
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main goroutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	log.Println("Register for signals.")
	signal.Notify(
		signalChannel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	<-signalChannel
	log.Println("Shutting  down...")

	go func() {
		<-signalChannel
		log.Fatalln("Terminating!")
	}()

	gracefulCtx, cancelShutdown := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefulCtx); err != nil {
		log.Printf("Shutdown error: %v\n", err)
		return err
	}

	log.Println("Server gracefully stopped.")

	return nil
}
