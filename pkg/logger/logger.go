package logger

import (
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	ModeConsole = "console"
	ModeFile    = "file"
)

var logger *Logger

func Info(msg string) {
	logger.Info(msg)
}

func Infof(msg string, args ...interface{}) {
	logger.Infof(msg, args...)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Warnf(msg string, args ...interface{}) {
	logger.Warnf(msg, args...)
}

func Error(msg string) {
	logger.Error(msg)
}

func Errorf(msg string, args ...interface{}) {
	logger.Errorf(msg, args...)
}

func Panic(msg string) {
	logger.Panic(msg)
}

func Panicf(msg string, args ...interface{}) {
	logger.Panicf(msg, args)
}

func Fatal(msg string) {
	logger.Fatal(msg)
}

func Fatalf(msg string, args ...interface{}) {
	logger.Fatalf(msg, args)
}

type Config struct {
	Mode       string `json:"mode"`
	Level      string `json:"level"`
	Filename   string `json:"filename" yaml:"filename"`
	MaxSize    int    `json:"max_size" yaml:"max_size"`
	MaxAge     int    `json:"max_age" yaml:"max_age"`
	Compress   bool   `json:"compress" yaml:"compress"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups"`
}

type Logger struct {
	*zap.SugaredLogger
	output io.Writer
}

func InitLogger(cfg Config) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.TimeEncoderOfLayout(time.DateTime)
	enc := zapcore.NewJSONEncoder(config)
	var output io.Writer
	switch cfg.Mode {
	case ModeConsole:
		output = os.Stdout
	case ModeFile:
		output = newLumberjack(cfg)
	default:
		output = os.Stdout
	}
	level := zapcore.InfoLevel
	switch cfg.Level {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
	case "error":
		level = zapcore.ErrorLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	}

	syncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(output))
	core := zapcore.NewCore(enc, syncer, level)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(zapLogger)
	logger = &Logger{
		SugaredLogger: zapLogger.Sugar(),
		output:        output,
	}
}

func GetOutput() io.Writer {
	return logger.output
}

type LumberjackOption struct {
	filename   string
	maxSize    int
	maxAge     int
	compress   bool
	maxBackups int
}

func defaultLumberjackOption() LumberjackOption {
	return LumberjackOption{
		filename:   "project.log",
		maxSize:    5,
		maxAge:     3,
		compress:   false,
		maxBackups: 0,
	}
}

func newLumberjack(cfg Config) io.Writer {
	defaultOpt := defaultLumberjackOption()
	if cfg.Filename != "" {
		defaultOpt.filename = cfg.Filename
	}
	if cfg.MaxSize != 0 {
		defaultOpt.maxSize = cfg.MaxSize
	}
	if cfg.MaxAge != 0 {
		defaultOpt.maxAge = cfg.MaxAge
	}
	defaultOpt.compress = cfg.Compress
	if cfg.MaxBackups != 0 {
		defaultOpt.maxBackups = cfg.MaxBackups
	}
	return &lumberjack.Logger{
		Filename:   defaultOpt.filename,
		MaxSize:    defaultOpt.maxSize,
		MaxAge:     defaultOpt.maxAge,
		Compress:   defaultOpt.compress,
		MaxBackups: defaultOpt.maxBackups,
		LocalTime:  true,
	}
}
