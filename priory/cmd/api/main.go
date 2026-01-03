// Copyright 2026 Uday Tiwari. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/udaycmd/Anask/priory/config"
	"github.com/udaycmd/Anask/priory/internal/server"
)

const (
	maxServerShutdownTimeout = 7 * time.Second
)

func shutdown(server *http.Server, done chan bool) {
	sigCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// wait for the signal
	<-sigCtx.Done()

	log.Println("Shutting down the server, press CTRL+C to force exit.")
	stop()

	toCtx, cancel := context.WithTimeout(context.Background(), maxServerShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(toCtx); err != nil {
		log.Printf("Server forced to stop with an error: %v\n", err)
	}

	log.Println("Exiting the shutdown service.")

	// notify the main goroutine about the shutdown completion
	done <- true
}

func main() {
	cfg, err := config.Init("")
	if err != nil {
		log.Printf("Unable to load the config: %v\n", err)
	}
	
	// set gin release mode
	if cfg.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} 

	done := make(chan bool, 1)
	server := server.NewServer(cfg)

	go shutdown(server, done)

	log.Printf("Server started at port: %d", cfg.Priory.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Http server error: %v\n", err)
	}

	<-done
	log.Println("Server shutdown completed!")
}
