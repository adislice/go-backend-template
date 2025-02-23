package logger

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger inisiasi logger
func InitLogger() *zap.Logger {
	// Zap config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeCaller: relativeCallerEncoder,
	}

	// Console log config
	consoleEncoderConfig := encoderConfig
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	// File log config
	fileEncoderConfig := encoderConfig
	fileEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)

	// Buat file log
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Gagal membuka file log: %v", err)
	}

	// Setting zap untuk file log dan console log
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel), // Console (dengan warna)
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), zapcore.InfoLevel),       // File (tanpa warna)
	)

	// Inisiasi logger
	logger := zap.New(core, zap.AddCaller())

	// Set logger global variable
	Log = logger

	return logger
}

// relativeCallerEncoder menampilkan nama file dan line number pada log
func relativeCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	baseProject, _ := os.Getwd()
	projectRoot := filepath.ToSlash(baseProject)
	relativePath := strings.TrimPrefix(caller.File, projectRoot)
	enc.AppendString(relativePath + ":" + strconv.Itoa(caller.Line))
}
