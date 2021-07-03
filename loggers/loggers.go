package loggers

import (
	"log"
	"os"
)

var Warning *log.Logger
var Error *log.Logger

func init(){
	warningFile, err := os.OpenFile("warning.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln("Couldn't open warning file")
	}

	errorFile, err := os.OpenFile("error.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln("Couldn't open error file")
	}

	Warning = log.New(errorFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(warningFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
