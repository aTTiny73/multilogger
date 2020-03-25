package multipleLogger

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
