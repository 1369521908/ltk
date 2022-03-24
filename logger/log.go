package logger

import "github.com/beego/beego/v2/core/logs"

func init() {
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterFile, `{"filename":"/logs/app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func Trace(f interface{}, v ...interface{}) {
	logs.Trace(f, v...)
}

func Debug(f interface{}, v ...interface{}) {
	logs.Debug(f, v...)
}

func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v...)
}

func Warn(f interface{}, v ...interface{}) {
	logs.Warn(f, v...)
}

func Error(f interface{}, v ...interface{}) {
	logs.Error(f, v...)
}

func Critical(f interface{}, v ...interface{}) {
	logs.Critical(f, v...)
}
