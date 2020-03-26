package logs

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

//DatabaseConfiguration setting up database
func DatabaseConfiguration() *sql.DB {
	conn, err := sql.Open("mysql", "TestUser:12345678@tcp/LOGGER")
	if err != nil {
		log.Print(err)
	}
	return conn
}

//NewDataBaseLog initilizes DataBaseLogger structure with DataBaseConfiguration
func NewDataBaseLog(DB *sql.DB) *DataBaseLogger {
	var DBLog DataBaseLogger
	DBLog.database = DB
	return &DBLog
}

//WriteToDB writes to database
func (dblog *DataBaseLogger) WriteToDB(str string) {
	stmt, err := dblog.database.Prepare("INSERT INTO Log(PREFIX, DATE, TIME, MESSAGE) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
	}
	date := fmt.Sprint(time.Now().Format("01-02-2006"))
	time := fmt.Sprint(time.Now().Format("15:04:05"))
	_, err = stmt.Exec(dblog.prefix, date, time, str)
	if err != nil {
		log.Print(err)
	}
}

//SetPrefix sets the log prefix
func (dblog *DataBaseLogger) SetPrefix(prefix string) {
	dblog.prefix = prefix
}

//Close the DataBase
func (dblog *DataBaseLogger) Close() {
	dblog.database.Close()
}

//Println for DataBase,converts the input to string and sends it to method WriteToDB
func (dblog *DataBaseLogger) Println(v ...interface{}) {
	dblog.WriteToDB(fmt.Sprint(v...))
}

//Printf for DataBase,converts the input to string and sends it to method WriteToDB
func (dblog *DataBaseLogger) Printf(format string, v ...interface{}) {
	dblog.WriteToDB(fmt.Sprintf(format, v...))
}
