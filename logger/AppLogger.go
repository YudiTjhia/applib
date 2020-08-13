package logger

import (
	"applib/conf"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)

func logMessage(logID string, msg string, tag string) {
	f1 := openFile()
	defer f1.Close()
	logMsg := logID + ">" + tag + msg
	printlnLog(logMsg)
}

func printlnLog(msg interface{}) {
	log.Println(msg)
	fmt.Println(msg)
}

func printLog(msg interface{}) {
	log.Print(msg)
	fmt.Print(msg)
}

func openFile() *os.File {

	now := time.Now()
	dt := now.Format("2006_01_02")

	appConf := conf.GetAppConf("") //need config file
	fileName := appConf.BasePath + "log/" + dt + ".log"

	f1, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f1)

	return f1
}

func logInterface(logID string, msg interface{}, tag string) {

	f1 := openFile()
	defer f1.Close()
	tag = logID + ">" + tag
	printlnLog(tag)
	printlnLog(msg)
}

func LogInfo(logID string, msg string) {
	logMessage(logID, msg, "[INFO]")
}

func LogError(logID string, msg string) {
	logMessage(logID, msg, "[ERROR]")
}

func LogBegin(serviceMethod string) string {
	logID := uuid.New().String()
	LogInfo(logID, "[BEGIN]="+serviceMethod)
	return logID
}

func LogGormSQL(logID string, query *gorm.DB) {
	sql, ok := query.Get("sql")
	if ok {
		LogInfo(logID, "[SQL]="+sql.(string))
	}

	sqlVars, ok := query.Get("sqlVars")
	if ok {
		b, _ := json.Marshal(sqlVars)
		logMessage(logID, string(b), "[SQLVARS]=")
	}

}

func LogParams(logID string, params map[string]string) {
	for k, v := range params {
		LogInfo(logID, "[PARAM]["+k+"]="+v)
	}
}

func LogData(logID string, data interface{}) {
	b, _ := json.Marshal(data)
	LogInfo(logID, "[Data]="+string(b))
}

func LogReturn(logID string, key string, value string) {
	LogInfo(logID, "[RETURN]["+key+"]="+value)
}

func LogReturns(logID string, returns map[string]string) {
	for k, v := range returns {
		LogInfo(logID, "[RETURN]["+k+"]="+v)
	}
}

func LogEnd(logID string) {
	LogInfo(logID, "[END]")
}
