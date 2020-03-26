package logs

import (
	"log"
	"log/syslog"
)

// NewSysLogger logs the system
func NewSysLogger(p syslog.Priority, f int) (*log.Logger, error) {
	var slog SysLogger
	slog.logger, slog.err = syslog.NewLogger(p, f)
	return slog.logger, slog.err
}

// Println function for SysLogger
func (slog *SysLogger) Println(v ...interface{}) {
	slog.logger.Println(v...)
}

// Printf function for SysLogger
func (slog *SysLogger) Printf(format string, v ...interface{}) {
	slog.logger.Printf(format, v...)
}

// SetPrefix funcion for SysLoggeer
func (slog *SysLogger) SetPrefix(pref string) {
	slog.logger.SetPrefix(pref)
}
