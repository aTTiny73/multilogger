package logs

import (
	"fmt"
	"log"
	"os"

	"log/syslog"

	"database/sql"
)

//LogFace interface
type LogFace interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	SetPrefix(s string)
}

////////////////////////////////////////////////////////////////////////////////

//FileLogger struct
type FileLogger struct {
	fd     *os.File
	logger *log.Logger
}

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

///////////////////////////////////////////////////////////////////////

//StdLogger strucure
type StdLogger struct {
	logger *log.Logger
}

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

////////////////////////////////////////////////////////////////////////

//SysLogger struct
type SysLogger struct {
	logger *log.Logger
	err    error
}

//NewSysLogger logs the system
func NewSysLogger(p syslog.Priority, f int) (*log.Logger, error) {
	var slog SysLogger
	slog.logger, slog.err = syslog.NewLogger(p, f)
	return slog.logger, slog.err
}

//Println function for SysLogger
func (slog SysLogger) Println(v ...interface{}) {
	slog.logger.Println(v...)
}

//Printf function for SysLogger
func (slog SysLogger) Printf(format string, v ...interface{}) {
	slog.logger.Printf(format, v...)
}

//SetPrefix funcion for SysLoggeer
func (slog SysLogger) SetPrefix(pref string) {
	slog.logger.SetPrefix(pref)
}

//////////////////////////////////////////////////////////////////////////

//DataBaseLogger struct
type DataBaseLogger struct {
	database *sql.DB
	prefix   string
}

//////////////////////////////////////////////////////////////////////////

//MultipleLog struct for multiplelogs
type MultipleLog struct {
	sliceLog []LogFace
	flag     bool
}

//NewCustomLogger makes a slice of loggers
func NewCustomLogger(flag bool, logs ...LogFace) MultipleLog {
	var multilog MultipleLog
	multilog.sliceLog = logs
	multilog.flag = flag
	return multilog
}

//Info method
func (multilog MultipleLog) Info(v ...interface{}) {
	for _, logg := range multilog.sliceLog {
		logg.SetPrefix("Info")
		logg.Println(v)
	}
}

//Infof method
func (multilog MultipleLog) Infof(format string, v ...interface{}) {
	for _, logg := range multilog.sliceLog {
		logg.SetPrefix("Infof")
		logg.Printf(format, v)
	}
}

//Warn method
func (multilog MultipleLog) Warn(v ...interface{}) {
	for _, logg := range multilog.sliceLog {
		logg.SetPrefix("Warn")
		logg.Println(v)
	}
}

//Warnf method
func (multilog MultipleLog) Warnf(format string, v ...interface{}) {
	for _, logg := range multilog.sliceLog {
		logg.SetPrefix("Warnf")
		logg.Printf(format, v)
	}
}

//Error method
func (multilog MultipleLog) Error(v ...interface{}) {
	for _, logg := range multilog.sliceLog {
		logg.SetPrefix("Error")
		logg.Println(v)
	}
}

//Errorf method
func (multilog MultipleLog) Errorf(format string, v ...interface{}) {
	for _, logg := range multilog.sliceLog {
		logg.SetPrefix("Errorf")
		logg.Printf(format, v)
	}
}

//Debug method
func (multilog MultipleLog) Debug(v ...interface{}) {
	if multilog.flag {
		for _, logg := range multilog.sliceLog {
			logg.SetPrefix("Debug")
			logg.Println(v)
		}
	}
}

//Debugf method
func (multilog MultipleLog) Debugf(format string, v ...interface{}) {
	if multilog.flag {
		for _, logg := range multilog.sliceLog {
			logg.SetPrefix("Debugf")
			logg.Printf(format, v)
		}
	}
}

/*CREATE TABLE Log
(
    NUM int NOT NULL AUTO_INCREMENT,
    PREFIX varchar(255) NOT NULL,
    DATE varchar(255) NOT NULL,
    TIME varchar(255) NOT NULL,
    MESSAGE varchar(255) NOT NULL,
    PRIMARY KEY (NUM)
);*/
