package stdLogger

import (
	"log"
	"os"
)

//NewStdLogger creats new std out logger
func NewStdLogger() *StdLogger {
	var stdlog StdLogger
	stdlog.logger = log.New(os.Stdout, "? ", log.LstdFlags)
	//log.SetOutput(os.Stdout)
	return &stdlog
}

//SetPrefix funcion
func (stdlog StdLogger) SetPrefix(pref string) {
	stdlog.logger.SetPrefix(pref)
}

//Println function for StdLogger
func (stdlog StdLogger) Println(v ...interface{}) {
	stdlog.logger.Println(v...)
}

//Printf function for StdLogger
func (stdlog StdLogger) Printf(format string, v ...interface{}) {
	stdlog.logger.Printf(format, v...)
}
