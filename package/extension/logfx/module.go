package logfx

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(time.RFC3339))
}

type Params struct {
	fx.In
	Lever string `name:"lever"`
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

func ForRoot(logLevel string) fx.Option {
	return fx.Module("gestlog",
		fx.Provide(fx.Annotate(
			func() string {
				return logLevel
			},
			fx.ResultTags(`name:"lever"`))),
		fx.Provide(ProvideLogger),
	)
}

func ContextDetail(ctx context.Context, keys ...string) string {
	var sb strings.Builder
	sb.WriteString("[")

	var parts []string
	for _, key := range keys {
		value := ctx.Value(key)
		parts = append(parts, fmt.Sprintf("%s: %v", key, value))
	}

	sb.WriteString(strings.Join(parts, ", "))
	sb.WriteString("]")
	return sb.String()
}
