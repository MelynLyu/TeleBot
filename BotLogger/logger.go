package BotLogger

import (
	"log"
	"os"
)

type BotLogger struct {
	logger *log.Logger
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func (bl *BotLogger) Setup(logFile string) {
	var file *os.File
	var err error

	if logFile == "" {
		file = os.Stderr
	} else {
		basePath := "~/logs/TeleBot/"
		fileName := basePath + logFile
		file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
	}
	bl.logger = log.New(file, "", log.LstdFlags)
	bl.logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func (bl BotLogger) Output(v ...interface{}) {
	bl.logger.Print(v...)
}

func (bl BotLogger) Outputf(format string, v ...interface{}) {
	bl.logger.Printf(format, v...)
}

func (bl BotLogger) Outputln(v ...interface{}) {
	bl.logger.Println(v...)
}
