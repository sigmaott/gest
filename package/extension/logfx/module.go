package logfx

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}

type Params struct {
	fx.In
	Lever string `optional:"true"`
}

func GetLogLever(level string) zapcore.Level {
	var logLevelSeverity = map[string]zapcore.Level{
		"DEBUG":     zapcore.DebugLevel,
		"INFO":      zapcore.InfoLevel,
		"WARNING":   zapcore.WarnLevel,
		"ERROR":     zapcore.ErrorLevel,
		"CRITICAL":  zapcore.DPanicLevel,
		"ALERT":     zapcore.PanicLevel,
		"EMERGENCY": zapcore.FatalLevel,
	}
	if zLever, ok := logLevelSeverity[level]; !ok {
		return zapcore.DebugLevel
	} else {
		return zLever
	}
}
func ProvideLogger(params Params) *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = SyslogTimeEncoder
	config.DisableStacktrace = true
	config.Encoding = "console"

	config.Level.SetLevel(zapcore.Level(GetLogLever(params.Lever)))

	logger, _ := config.Build()

	return logger.Sugar()
}

//// Module provided to fx
//var Module = fx.Module("gestlog",
//	fx.Provide(ProvideLogger),
//)

func Module() fx.Option {
	return fx.Module("gestlog",
		fx.Provide(ProvideLogger),
	)
}
