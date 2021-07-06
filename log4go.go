package logger

import (
	"fmt"
	"go.uber.org/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"strconv"
	"strings"
	"time"
)

var sugar *zap.SugaredLogger
var Logger *zap.Logger

// getLumberjackLogger 获取lumberjack.Logger
func getLumberjackLogger(root *config.YAML, level string) *lumberjack.Logger {
	option := root.Get("LOG4GO").Get(level)
	filename := option.Get("FILE_PATH_NAME").String()
	maxSize, err := strconv.Atoi(option.Get("MAXSIZE").String())
	if err != nil {
		log.Fatalln(fmt.Sprintf("%s MAXSIZE参数不为数字或者参数不合法", level))
	}
	maxBackups, err := strconv.Atoi(option.Get("MAXBACKUP_COUNT").String())
	if err != nil {
		log.Fatalln(fmt.Sprintf("%s MAXBACKUP_COUNT参数不为数字或者参数不合法", level))
	}
	maxAge, err := strconv.Atoi(option.Get("MAXAGE").String())
	if err != nil {
		log.Fatalln(fmt.Sprintf("%s MAXAGE参数不为数字或者参数不合法", level))
	}
	compress, err := strconv.ParseBool(option.Get("COMPRESS").String())
	if err != nil {
		log.Fatalln(fmt.Sprintf("%s COMPRESS参数不为数字或者参数不合法", level))
	}
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge,   // days
		Compress:   compress, // disabled by default
	}
}

func init() {
	var options config.YAMLOption = config.File("log4go.yml")
	root, err := config.NewYAML(options)
	if err != nil {
		log.Fatalln(err)
	}
	log4goFormat := root.Get("LOG4GO").Get("FORMAT").String()
	log4goLevelMode := root.Get("LOG4GO").Get("LEVEL_MODE").String()
	log4goLevelColor, err := strconv.ParseBool(root.Get("LOG4GO").Get("LEVEL_COLOR").String())
	if err != nil {
		log.Fatalln("LEVEL_COLOR 参数不合法")
	}

	var encode zapcore.Encoder
	e := zap.NewProductionEncoderConfig()
	e.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006.01.02 15:04:05"))
	}
	if log4goLevelColor {
		e.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		e.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	// 判断是Json配置还是Text配置
	if strings.ToUpper(log4goFormat) == "TEXT" {
		encode = zapcore.NewConsoleEncoder(e)
	} else if strings.ToUpper(log4goFormat) == "JSON" {
		encode = zapcore.NewJSONEncoder(e)
	} else {
		encode = zapcore.NewJSONEncoder(e)
	}

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		switch log4goLevelMode {
		case "contain":
			return lvl >= zapcore.InfoLevel
		case "independent":
			return lvl >= zapcore.InfoLevel && lvl < zapcore.ErrorLevel
		default:
			return lvl >= zapcore.InfoLevel
		}
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	hookInfo := getLumberjackLogger(root, "INFO")
	defer func(hookInfo *lumberjack.Logger) {
		if err := hookInfo.Close(); err != nil {
			log.Fatalln(err)
		}
	}(hookInfo)

	hookError := getLumberjackLogger(root, "ERROR")
	defer func(hookError *lumberjack.Logger) {
		if err := hookError.Close(); err != nil {
			log.Fatalln(err)
		}
	}(hookError)

	core := zapcore.NewTee(
		zapcore.NewCore(encode, zapcore.AddSync(hookInfo), infoLevel),
		zapcore.NewCore(encode, zapcore.AddSync(hookError), errorLevel),
	)
	Logger = zap.New(core, zap.AddCaller())
	defer func(logger *zap.Logger) {
		if err := logger.Sync(); err != nil {
			log.Fatalln(err)
		}
	}(Logger)
	sugar = Logger.Sugar()
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Infof(temp string, args ...interface{}) {
	sugar.Infof(temp, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Warnf(temp string, args ...interface{}) {
	sugar.Warnf(temp, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}

func Errorf(temp string, args ...interface{}) {
	sugar.Errorf(temp, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	sugar.Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	sugar.Fatal(args...)
}

func Fatalf(temp string, args ...interface{}) {
	sugar.Fatalf(temp, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}
