package logs

import (
	"fmt"
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
	return &flog
}

//Println function for FileLogger
func (flog *FileLogger) Println(v ...interface{}) {
	date, time := getDateTime()
	text := fmt.Sprintln(v...)
	output := []byte(flog.prefix + " " + date + " " + time + " " + text)
	flog.fd.Write(output)
}

//Printf function for FileLogger
func (flog *FileLogger) Printf(format string, v ...interface{}) {
	date, time := getDateTime()
	text := fmt.Sprintf(format, v...)
	output := []byte(flog.prefix + " " + date + " " + time + " " + text + "\n")
	flog.fd.Write(output)
}

//Close for FileLogger
func (flog *FileLogger) Close() {
	flog.fd.Close()
}

//SetPrefix funcion
func (flog *FileLogger) SetPrefix(pref string) {
	flog.prefix = pref
}
