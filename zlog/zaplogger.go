package zlog

import (
	"flag"
	"fmt"
	"go.uber.org/zap/zapcore"
	"runtime"
	"time"
)

type ZapLogger interface {
	Error(string)
	Errorf(string, ...interface{})
	Info(string)
	Infof(string, ...interface{})
	Fatal(string)
}

type ZapConfig struct {
	// Logger类型 json console 两种
	LoggerType string
	// 服务名称
	ServerName string
	// 日志级别
	Level zapcore.Level
	// 日志存放目录
	LogDir string
	// 日志切割大小 单位 Mb
	CutSize int
	// 是否压缩
	Compress bool

	// 企业微信通知 默认通知
	Alarm bool
	// 企业微信通知人 多个逗号分割
	AlarmUser string
	// 企业微信通知组 多个逗号分割
	AlarmGroup string
	// mention
	Mention string

	//当前的logger
	logger ZapLogger
}

var (
	zapCfg  *ZapConfig
	log_dir = "./"
)

func init() {
	flag.StringVar(&log_dir, "log_dir", log_dir, "setting log dir")
}

// 初始zap配置
func NewZapConfig() *ZapConfig {
	// 默认配置
	zapCfg = &ZapConfig{
		LoggerType: "console",
		ServerName: "UNKNOW",
		Level:      zapcore.DebugLevel,
		LogDir:     log_dir,
		CutSize:    3,
		Compress:   false,
	}
	return zapCfg
}

// 创建zap logger
func (z *ZapConfig) Build() {
	if z.LoggerType == "json" {
		z.logger = z.JsonLogger()
	} else {
		z.logger = z.ConsoleLogger()
	}
}

func Errorf(str string, arg ...interface{}) {
	zapCfg.logger.Errorf(str, arg...)
}

func Error(str string) {
	zapCfg.logger.Error(str)
}

func Infof(str string, arg ...interface{}) {
	zapCfg.logger.Infof(str, arg...)
}

func Info(str string) {
	zapCfg.logger.Info(str)
}

func Fatal(str string) {
	zapCfg.logger.Fatal(str)
}

// 自定义时间格式
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 跳过指定的栈信息
func getFuncStack(skip int) string {
	funcName, _, line, ok := runtime.Caller(skip)
	if ok {
		return fmt.Sprintf("%s[%d]", runtime.FuncForPC(funcName).Name(), line)
	}
	return ""
}

// Error消息告警
func (z *ZapConfig) wxwrokAlarm(errStr string) {

}
