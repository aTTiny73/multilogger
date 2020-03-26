package main

import (
	"log/syslog"

	"github.com/aTTiny73/multilogger/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mkmueller/golog"
)

func main() {

	fileLog1 := logs.NewFileLogger("fileLog1")
	defer fileLog1.Close()
	//fileLog1.SetPrefix("  Info")
	//fileLog1.Println("Test file log")

	fileLog2 := logs.NewFileLogger("fileLog2")
	defer fileLog2.Close()
	//fileLog2.Println("Test filelog2")

	syslog, _ := logs.NewSysLogger(syslog.LOG_NOTICE, golog.LstdFlags)
	//syslog.SetPrefix("SysLOG")
	//syslog.Println("sylog test")

	stdlog := logs.NewStdLogger()
	defer stdlog.Close()
	//stdlog.SetPrefix("Test")
	//stdlog.Println("StdLog Test")

	databaseLog := logs.NewDataBaseLog(logs.DatabaseConfiguration())
	//databaseLog.SetPrefix("INFO")
	//databaseLog.WriteToDB("database prefix test")

	log := logs.NewCustomLogger(true, fileLog1, fileLog2, syslog, databaseLog)

	log.Info("info")
	log.Infof("%s", "infof")
	log.Warn("warn")
	log.Warnf("%s", "warnf")
	log.Debug("debug")
	log.Debugf("%s", "debugf")
	log.Error("error")
	log.Errorf("%s", "errorf")
}
