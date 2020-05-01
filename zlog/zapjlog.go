package zlog

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
)

// 日志输出格式为Json

type jsonLogger struct {
	log *zap.Logger
}

func(z *ZapConfig) JsonLogger() *jsonLogger{
	fileName := fmt.Sprintf("%s.json.log", z.ServerName)
	logPath := path.Join(z.LogDir, fileName)
	var jlog jsonLogger
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "Time",
		LevelKey:       "LV",
		NameKey:        "Name",
		MessageKey:     "Msg",
		CallerKey:      "Caller",
		EncodeTime:     TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	hook := lumberjack.Logger{
		Filename: logPath,   // 日志文件路径
		MaxSize:  z.CutSize, // megabytes
		Compress: false,     // 是否压缩 disabled by default
	}
	atom := zap.NewAtomicLevelAt(z.Level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atom,
	)
	jlog.log = zap.New(
		core,
		zap.AddCaller(),
		zap.Development(),
		zap.AddCallerSkip(2),   // 指定跳过堆栈，因为做了封装，需要返回实际抛错的函数
	).Named(z.ServerName)

	return &jlog
}

func(jl *jsonLogger) Errorf(str string, arg ...interface{}){
	var s = str
	if len(arg) != 0{
		s = fmt.Sprintf(str, arg...)
	}
	zapCfg.wxwrokAlarm(s)
	jl.log.Error(s)
}

func(jl *jsonLogger) Error(str string){
	zapCfg.wxwrokAlarm(str)
	jl.log.Error(str)
}

func(jl *jsonLogger) Infof(str string, arg ...interface{}){
	var s = str
	if len(arg) != 0{
		s = fmt.Sprintf(str, arg...)
	}
	jl.log.Info(s)
}

func(jl *jsonLogger) Info(str string){
	jl.log.Info(str)
}

func(jl *jsonLogger) Fatal(str string){
	zapCfg.wxwrokAlarm(str)
	jl.log.Fatal(str)
}
