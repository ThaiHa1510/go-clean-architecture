package logs

import(
	"time"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
) 

func LoadLogger() *zap.SugaredLogger{
	return ConfigZap()
}

