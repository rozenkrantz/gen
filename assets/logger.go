package assets

func GetLogger() string {
	return `package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//Module is one main modules
var Module = fx.Provide(NewLogger)

//NewLogger create new Logger type
func NewLogger() (logger *zap.Logger, err error) {

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "log/app.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     10,
	})

	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "datetime"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.DebugLevel,
	)

	logger = zap.New(core)

	return logger, nil

}

func Log(logger *zap.Logger, message, service, function, operation string) {
	logger.Error(message,
		zap.String("Service", service),
		zap.String("Function", function),
		zap.String("Operation", operation),
	)
}
`
}
