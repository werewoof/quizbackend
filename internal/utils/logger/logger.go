package logger

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	WarnLogger  *log.Logger
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
	FatalLogger *log.Logger
	DebugLogger *log.Logger
)

//replace with Zap library later

func init() {
	//get current date for file
	t := time.Now()
	name := t.Format("2006-01-02") + ".log"
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	//print to file and console
	output := io.MultiWriter(os.Stdout, file)
	errOutput := io.MultiWriter(os.Stderr, file)

	InfoLogger = log.New(output, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarnLogger = log.New(errOutput, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errOutput, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(errOutput, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(output, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
