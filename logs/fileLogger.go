package logs

import (
	"fmt"
	"log"
	"os"
)

//NewFileLogger creats a new file and returns a pointer to that file
func NewFileLogger(s string) *FileLogger {
	var flog FileLogger
	file, err := os.OpenFile(s, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Can not open file")
	}
	flog.fd = file
	flog.logger = log.New(file, "? ", log.LstdFlags)
	return &flog
}

//Println function for FileLogger
func (flog FileLogger) Println(v ...interface{}) {
	flog.logger.Println(v...)
}

//Printf function for FileLogger
func (flog FileLogger) Printf(format string, v ...interface{}) {
	flog.logger.Printf(format, v...)
}

//Close for FileLogger
func (flog FileLogger) Close() {
	flog.fd.Close()
}

//SetPrefix funcion
func (flog FileLogger) SetPrefix(pref string) {
	flog.logger.SetPrefix(pref)
}
