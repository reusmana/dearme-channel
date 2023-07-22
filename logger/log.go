package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rals/dearme-channel/enums"
)

// GeneralLogger exported
var GeneralLogger *log.Logger

// ErrorLogger exported
var ErrorLogger *log.Logger

// GeneralLogger exported
var InfoLogger *log.Logger

// GeneralLogger exported
var WarningLogger *log.Logger

func init() {
	currentTime := time.Now()
	absPath, err := filepath.Abs(enums.PATH_LOCAL + "/log")
	// absPath, err := filepath.Abs(enums.PATH_SERVER + "/log")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/general-log-"+currentTime.Format("2006-01-02")+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	GeneralLogger = log.New(generalLog, "General Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(generalLog, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(generalLog, "Info Logger: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(generalLog, "Warning Logger: ", log.Ldate|log.Ltime|log.Lshortfile)
}
