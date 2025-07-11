package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/myrat012/to-do/internal/controller/http"
	"github.com/myrat012/to-do/internal/usecase"
	"github.com/myrat012/to-do/pkg/config"
)

func main() {
	cfg, err := config.ReadConfig("config/config.yaml")
	if err != nil {
		log.Printf("Error loading config: %v\n", err)
		return
	}

	// Use cases
	useCases := usecase.LoadUseCases()

	// Waiting signal
	signalChan := make(chan os.Signal, 1)
	quit := make(chan interface{})

	srv := http.NewService(cfg, useCases)
	if err != nil {
		fmt.Printf("could not initialize http.NewService: %v", err)
		return
	}

	var listener net.Listener
	listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Printf("error in setting up listener: %v", err)
		return
	}

	go func() {
		err = srv.Server.Serve(listener)
		if err != nil {
			log.Printf("HTTP server Error: %v", err)
		}
	}()

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-quit:
			log.Printf("quit channel closed, closing listener: %v", err)
			err = srv.Server.Close()
			if err != nil {
				log.Printf("error during HTTP Server close: %v", err)
			}
			err = listener.Close()
			if err != nil {
				log.Printf("error during TCP Listener close: %v", err)
			}
			return
		case sig := <-signalChan:
			switch sig {
			case os.Interrupt, os.Kill, syscall.SIGTERM:
				log.Printf("interrupt signal received, sending Quit signal: %v", sig)
				close(quit)
			default:
				log.Printf("signal received: %v", sig)
			}
		}
	}
}
