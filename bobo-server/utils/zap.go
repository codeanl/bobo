package utils

import (
	"bobo-server/config"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"time"
)

// Logger 全局日志指针
var Logger *zap.Logger

func InitLogger() {
	// 生成日志文件目录
	if ok, _ := PathExists(config.Cfg.Zap.Directory); !ok {
		_ = os.Mkdir(config.Cfg.Zap.Directory, os.ModePerm)
	}
	core := zapcore.NewCore(getEncoder(), getWriterSyncer(), getLevelPriority())
	Logger = zap.New(core)
	if config.Cfg.Zap.ShowLine {
		// 获取 调用的文件, 函数名称, 行号
		Logger = Logger.WithOptions(zap.AddCaller())
	}
	log.Println("ZapLogger初始化成功")
}

// 编码器: 如何写入日志
func getEncoder() zapcore.Encoder {
	// 参考: zap.NewProductionEncoderConfig()
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // ?
	}
	if config.Cfg.Zap.Format == "json" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志输出路径: 文件、控制台、双向输出
func getWriterSyncer() zapcore.WriteSyncer {
	file, _ := os.Create(config.Cfg.Zap.Directory + "/log.log")
	// 双向输出
	if config.Cfg.Zap.LogInConsole {
		fileWriter := zapcore.AddSync(file)
		consoleWriter := zapcore.AddSync(os.Stdout)
		return zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter)
	}
	// 输出到文件
	return zapcore.AddSync(file)
}

// 获取日志输出级别
func getLevelPriority() zapcore.LevelEnabler {
	switch config.Cfg.Zap.Level {
	case "debug", "Debug":
		return zap.DebugLevel
	case "info", "Info":
		return zap.InfoLevel
	case "warn", "Warn":
		return zap.WarnLevel
	case "error", "Error":
		return zap.ErrorLevel
	case "dpanic", "DPanic":
		return zap.DPanicLevel
	case "panic", "Panic":
		return zap.PanicLevel
	case "fatal", "Fatal":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(config.Cfg.Zap.Prefix + t.Format("2006/01/02 - 15:04:05"))
}

// PathExists 判断文件目录是否存在
func PathExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	// 存在文件夹
	if fileInfo.IsDir() {
		return true, nil
	}
	// 不是文件夹, 则存在同名文件
	return false, errors.New("存在同名文件")
}
