package logs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func SetupLogger(ctx context.Context) *os.File {
	logPath := filepath.Join("logs", "chat.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file")
		Log.Fatal(err)
		return nil
	}
	defer file.Close()
	Log.SetOutput(file)
	Log.SetFormatter(&logrus.JSONFormatter{})
	return file
}