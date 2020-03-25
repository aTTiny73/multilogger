package logs

import (
	"log"
	"os"

	"database/sql"
)

//LogFace interface
type LogFace interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	SetPrefix(s string)
}

//FileLogger struct
type FileLogger struct {
	fd     *os.File
	logger *log.Logger
}

//StdLogger strucure
type StdLogger struct {
	logger *log.Logger
}

//SysLogger struct
type SysLogger struct {
	logger *log.Logger
	err    error
}

//DataBaseLogger struct
type DataBaseLogger struct {
	database *sql.DB
	prefix   string
}

//MultipleLog struct for multiplelogs
type MultipleLog struct {
	sliceLog []LogFace
	flag     bool
}
