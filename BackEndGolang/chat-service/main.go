package main

import (
	"chat-config/config"
	"chat-logs/logs"
	"chat-service/cmd"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	// Set output to a file
	logPath := filepath.Join("internal", "logs", "chat.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file")
		logs.Log.Fatal(err)
		return
	}
	defer file.Close()
	logs.Log.SetOutput(file)
	logs.Log.SetFormatter(&logrus.JSONFormatter{})
	configPath := filepath.Join("config", "config.yaml")
	config, err := config.LoadConfig(configPath)

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
