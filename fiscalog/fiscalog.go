package fiscalog

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	currPath, err := os.Getwd()

	if err != nil {
		log.Fatalln("Failed to get current dir: ", err)
	}

	logFilePath := currPath + "/fiscalog/fiscaluno.log"
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to open log file: ", err)
	}

	Trace = log.New(io.MultiWriter(logFile, ioutil.Discard),
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(logFile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}
