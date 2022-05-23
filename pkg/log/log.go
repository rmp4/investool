package log

import (
	"investool/pkg/utils"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once           sync.Once
	log            *zap.Logger
	processLogLock sync.Once
	processLog     *zap.Logger
)

func ServiceLog() *zap.Logger {
	if log == nil {
		once.Do(func() {
			path, _ := utils.GetCurrentPath()
			logFile := filepath.Clean(path + "/service.log")
			writeFile := zapcore.AddSync(&lumberjack.Logger{
				Filename: logFile,
				MaxSize:  10, // megabytes
				MaxAge:   28, // days
			})
			writeStdout := zapcore.AddSync(os.Stdout)
			encoderCfg := zap.NewProductionEncoderConfig()
			encoderCfg.TimeKey = "timestamp"
			encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
			core := zapcore.NewCore(
				zapcore.NewJSONEncoder(encoderCfg),
				zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
				zap.InfoLevel,
			)
			log = zap.New(
				core,
				zap.AddCaller(),
				zap.AddStacktrace(zap.ErrorLevel),
			)
		})
	}
	return log
}
