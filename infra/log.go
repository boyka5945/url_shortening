package infra

import (
	"fmt"
	"log"
	"os"
)

var (
	// GlobalLogger is the instance of the global logger
	logger *log.Logger
)

func GetLogger() *log.Logger {
	return logger
}

func InitLog() error {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %s", err)
	}

	// Create a logger instance
	logger = log.New(file, "", log.Ldate|log.Ltime)
	return nil
}
