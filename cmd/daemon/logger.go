package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger sets up a standardized logger according to standards setup by the search.
// Based on the `appenv` argument, the debugging level will be set to Info ("production") or Debug (anything else).
func logger(appenv string) (*zap.Logger, error) {
	logConf := zap.NewProductionConfig()

	switch appenv {
	case "production", "live":
		logConf.Level.SetLevel(zap.InfoLevel) // be explicit.
	default:
		logConf.Level.SetLevel(zap.DebugLevel)
	}

	logConf.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "message",
		LineEnding: zapcore.DefaultLineEnding,

		TimeKey:        "timestamp",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		StacktraceKey: "stacktrace",

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	logger, err := logConf.Build()
	if err != nil {
		return nil, err
	}

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)

	return logger, nil
}
