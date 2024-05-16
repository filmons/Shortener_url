// package utils

// import (
// 	"log"
// 	"os"
// )

// // Logger is the application-wide logger.
// var Logger *log.Logger

// func init() {
// 	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
// }

package utils

import (
    "log"
    "os"
)

var (
    // Logger is the application-wide logger for standard output.
    Logger *log.Logger
    // FileLogger is a logger instance for writing logs to files.
    FileLogger *log.Logger
)

func init() {
    Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

    // Open log file for writing, create if not exists, append if already exists
    file, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    FileLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs a message at INFO level to both standard output and log file.
func Info(message string) {
    Logger.Println(message)
    FileLogger.Println("INFO: " + message)
}

// Warn logs a message at WARN level to both standard output and log file.
func Warn(message string) {
    Logger.Println("WARN: " + message)
    FileLogger.Println("WARN: " + message)
}

// Error logs a message at ERROR level to both standard output and log file.
func Error(message string) {
    Logger.Println("ERROR: " + message)
    FileLogger.Println("ERROR: " + message)
}

