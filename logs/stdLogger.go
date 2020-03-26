package logs

import (
	"fmt"
	"os"
)

// NewStdLogger creats new std out logger
func NewStdLogger() *StdLogger {
	var stdlogger StdLogger
	file, err := os.OpenFile("/dev/stdout", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	stdlogger.fd = file
	return &stdlogger
}

// Println function for StdLogger
func (stdlog *StdLogger) Println(v ...interface{}) {
	date, time := getDateTime()
	text := fmt.Sprintln(v...)
	output := []byte(stdlog.prefix + " " + date + " " + time + " " + text)
	stdlog.fd.Write(output)
}

// Printf function for StdLogger
func (stdlog *StdLogger) Printf(format string, v ...interface{}) {
	date, time := getDateTime()
	text := fmt.Sprintf(format, v...)
	output := []byte(stdlog.prefix + " " + date + " " + time + " " + text)
	stdlog.fd.Write(output)
}

// Close closes the /dev/stdout
func (stdlog *StdLogger) Close() {
	stdlog.fd.Close()
}

// SetPrefix funcion for StdLogger
func (stdlog *StdLogger) SetPrefix(pref string) {
	stdlog.prefix = pref
}
