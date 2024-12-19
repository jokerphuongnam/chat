package main

import (
	"chat-config/config"
	"chat-logs/logs"
	"chat-service/cmd"
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Set output to a file
	file := logs.SetupLogger(context.Background())
	if file != nil {
		defer file.Close()
	}
	
	configPath := filepath.Join("config", "config.yaml")
	config, err := config.LoadConfig(configPath, func(config string) string {
		return os.ExpandEnv(config)
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	r := cmd.Execute(config)
	if r == nil {
		fmt.Println("Error loading")
		return
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port), // Server port
		Handler:      r,                                                            // The Gin router
		ReadTimeout:  time.Duration(config.Server.ReadTimeout) * time.Second,       // Set read timeout (for reading request)
		WriteTimeout: time.Duration(config.Server.WriteTimeout) * time.Second,      // Set write timeout (for writing response)
		IdleTimeout:  time.Duration(config.Server.IdleTimeout) * time.Second,       // Set idle timeout
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logs.Log.Fatalf("Server error: %v", err)
	}
}
