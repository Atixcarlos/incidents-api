package utils

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Declare variables for server logs.
var (
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	LogDirectory string
)

// createLogFile creates a new log file for today's log.
func createLogFile(logType string) *os.File {

	now := time.Now()
	day := strconv.Itoa(now.Day())
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	subdirectory := "/" + year + "/" + month
	logfileName := LogDirectory + subdirectory + "/" + logType + "-" + day + ".log"

	os.MkdirAll(LogDirectory+subdirectory, 0755)

	f, err := os.OpenFile(logfileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Error opening file: %s %v \n", logfileName, err)
	}

	return f
}

// maintainLogRotation checks every 30 minutes if we need to switch to a new log file.
// Log files are maintained on a daily basis.
// We ignore time difference at daily log switchover due to 30 minute interval:
// The new log may be created up to 29 minutes after midnight.
func maintainLogRotation(myLog *log.Logger, myLogFile *os.File, logType string) {

	defer myLogFile.Close()

	now := time.Now()
	currentDay := strconv.Itoa(now.Day())

	for {

		select {

		case <-time.After(30 * time.Minute):

			hourLater := time.Now()
			dayLater := strconv.Itoa(hourLater.Day())

			// If the next day started, start a new log file.
			if dayLater != currentDay {
				currentDay = dayLater
				nextLogFile := createLogFile(logType)
				myLog.SetOutput(nextLogFile)
				myLogFile.Close()
				myLogFile = nextLogFile
			}
		}
	}
}

// InitLogs creates log files and maintains log file rotation.
func InitLogs(logDirectory string) {

	LogDirectory = logDirectory
	infoHandle := createLogFile("info")
	errorHandle := createLogFile("error")

	InfoLog = log.New(infoHandle, "INFO: ", log.LstdFlags)
	ErrorLog = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	go maintainLogRotation(ErrorLog, errorHandle, "error")
	go maintainLogRotation(InfoLog, infoHandle, "info")
}

// LogRequest logs an API hit in the server logs.
func LogRequest(R *http.Request) {
	InfoLog.Printf("\t%s\t%s\t%s\t%s\n", R.Method, R.RequestURI, R.RemoteAddr, R.UserAgent())
}

// LogError logs an error in the server logs.
func LogError(R *http.Request, err error) {
	ErrorLog.Printf("\n%s\t%s\t%s\t%s\n%s\n", R.Method, R.RequestURI, R.RemoteAddr, R.UserAgent(), err.Error())
}
