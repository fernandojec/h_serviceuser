package loghelper

import (
	"context"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// app logger
var appLogger *zap.SugaredLogger

func init() {
	writerSyncer := getLogWriter()
	fileEncoder, consoleEncoder := getEncoder()
	defaultLogLevel := zap.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writerSyncer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	appLogger = logger.Sugar()
	defer appLogger.Sync()
}

func getEncoder() (zapcore.Encoder, zapcore.Encoder) {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	return fileEncoder, consoleEncoder //zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./storage/logs/app.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(ctx context.Context, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Debug(args...)
}

func Debugf(ctx context.Context, template string, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Debugf("%v "+template, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Info(args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Infof("%v "+template, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Warn(args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Warnf("%v "+template, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Error(args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Errorf("%v "+template, args...)
}

func DPanic(ctx context.Context, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.DPanic(args...)
}

func DPanicf(ctx context.Context, template string, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.DPanicf("%v "+template, args...)
}

func Panic(ctx context.Context, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Panic(args...)
}

func Panicf(ctx context.Context, template string, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Panicf("%v "+template, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Fatal(args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	args = appendContext(ctx, args...)
	appLogger.Fatalf("%v "+template, args...)
}

func appendContext(ctx context.Context, args ...interface{}) []interface{} {
	jsonData := map[string]interface{}{}
	xTraceID := ctx.Value(XTRACEID).(string)
	if xTraceID != "" {
		jsonData["TRACEID"] = xTraceID
		args = append([]interface{}{xTraceID}, args...)
	}
	jsonData["MESSAGE"] = args
	return args
}
