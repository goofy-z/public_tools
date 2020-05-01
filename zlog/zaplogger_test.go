package zlog

import (
	"testing"
)

// 功能测试
func TestLoggerFunction(t *testing.T){
	zcfg := NewZapConfig()
	zcfg.LoggerType = "console"
	zcfg.ServerName = "NEW-TEST"
	zcfg.Alarm = true
	zcfg.AlarmUser = "ginxzheng"
	zcfg.Build()
	Errorf("测试Errorf")
	Error("测试Error")
	Infof("测试Infof")
	Info("测试Info")
	Fatal("测试Fatal")
}
