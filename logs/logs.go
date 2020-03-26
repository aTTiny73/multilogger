package logs

import (
	"fmt"
	"log"
	"os"
	"time"

	"database/sql"
)

// LogFace interface
type LogFace interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	SetPrefix(s string)
}

// FileLogger struct
type FileLogger struct {
	fd     *os.File
	prefix string
}

// StdLogger strucure
type StdLogger struct {
	fd     *os.File
	prefix string
}

// SysLogger struct
type SysLogger struct {
	logger *log.Logger
	err    error
}

// DataBaseLogger struct
type DataBaseLogger struct {
	database *sql.DB
	prefix   string
}

// MultipleLog struct for multiplelogs
type MultipleLog struct {
	sliceLog []LogFace
	flag     bool
}

// Returns date(01-02-2006) and time(15:04:05)
func getDateTime() (string, string) {
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	return date, time
}
