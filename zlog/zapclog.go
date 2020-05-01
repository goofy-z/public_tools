package zlog

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
)

// 日志输出为Console形式

type consoleLogger struct {
	log *zap.Logger
}

func(z *ZapConfig) ConsoleLogger() *consoleLogger{
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "LV",
		NameKey:        "Name",
		MessageKey:     "Msg",
		CallerKey:      "Caller",
		// StacktraceKey:  "Stack", 关闭堆栈信息
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写编码
		EncodeTime:     TimeEncoder,                   // 时间编码
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     NewNameEncoder,
	}
	hook := lumberjack.Logger{
		Filename: path.Join(z.LogDir, fmt.Sprintf("%s.console.log", z.ServerName)), // 日志文件
		MaxSize:  z.CutSize,                // megabytes
		Compress: z.Compress,               // 是否压缩
	}
	atom := zap.NewAtomicLevelAt(z.Level)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atom,
	)
	var clog consoleLogger
	clog.log = zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(2)).Named(z.ServerName)
	return &clog
}

func(cl *consoleLogger) Errorf(str string, arg ...interface{}){
	var s = str
	if len(arg) != 0{s = fmt.Sprintf(str, arg...)}
	zapCfg.wxwrokAlarm(s)
	cl.log.Error(s)
}

func(cl *consoleLogger) Error(str string){
	zapCfg.wxwrokAlarm(str)
	cl.log.Error(str)
}

func(cl *consoleLogger) Infof(str string, arg ...interface{}){
	var s = str
	if len(arg) != 0{s = fmt.Sprintf(str, arg...)}
	cl.log.Info(s)
}

func(cl *consoleLogger) Info(str string){
	cl.log.Info(str)
}

func(cl *consoleLogger) Fatal(str string){
	zapCfg.wxwrokAlarm(str)
	cl.log.Fatal(str, zap.Stack("stack"))
}

// 自定义consolelogger名称
func NewNameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	str := fmt.Sprintf("[server: %s]", loggerName)
	enc.AppendString(str)
}