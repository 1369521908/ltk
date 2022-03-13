package logger

import "github.com/beego/beego/v2/core/logs"

func init() {
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/app.logger","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func Trace(format string, v ...interface{}) {
	logs.Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	logs.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	logs.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	logs.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	logs.Error(format, v...)
}

func Critical(format string, v ...interface{}) {
	logs.Critical(format, v...)
}
