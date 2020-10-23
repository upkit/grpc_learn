package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/robfig/cron"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Conf struct {
	Path  string `yaml:"path"`  // 日志文件路径
	Level string `yaml:"level"` // 日志等级
}

var (
	loggerLevel zapcore.Level

	debugLogger  = new(zap.Logger)
	infoLogger   = new(zap.Logger)
	warnLogger   = new(zap.Logger)
	errorLogger  = new(zap.Logger)
	dpanicLogger = new(zap.Logger)
	panicLogger  = new(zap.Logger)
	fatalLogger  = new(zap.Logger)
)

var loggerMap = map[string]**zap.Logger{
	"debug":  &debugLogger,
	"info":   &infoLogger,
	"warn":   &warnLogger,
	"error":  &errorLogger,
	"dpanic": &dpanicLogger,
	"panic":  &panicLogger,
	"fatal":  &fatalLogger,
}

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func Init(cfg Conf) {
	if cfg.Path == "" {
		panic("undefined log path")
	}

	// 创建日志文件夹
	err := os.MkdirAll(filepath.Dir(cfg.Path), 0755)
	if err != nil {
		fmt.Printf("InitLogger create logDir failed, path:%s, err:%v\n", filepath.Base(cfg.Path), err)
		panic(err)
	}

	loggerLevel = getLoggerLevel(cfg.Level)
	for key, value := range levelMap {
		newLogger(key, cfg.Path, cfg.Level, value)
	}
}

func Close() {

}

func newLogger(key, file, level string, value zapcore.Level) {
	fileName := file + "_" + key + ".log"
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == value
	})

	var (
		console    zapcore.WriteSyncer
		syncWriter zapcore.WriteSyncer
	)

	l := &lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 7, //128m
		LocalTime: true,
		Compress:  true,
	}

	if value < getLoggerLevel(level) {
		syncWriter = zapcore.AddSync(ioutil.Discard)
		console = zapcore.Lock(syncWriter)
	} else {
		syncWriter = zapcore.AddSync(l)
		if value > zapcore.InfoLevel {
			console = zapcore.Lock(os.Stderr)
		} else {
			console = zapcore.Lock(os.Stdout)
		}
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = timeEncoder
	consoleEncoder := zap.NewDevelopmentEncoderConfig()
	consoleEncoder.EncodeTime = timeEncoder

	core := zapcore.NewTee(
		// 打印在控制台
		zapcore.NewCore(zapcore.NewConsoleEncoder(consoleEncoder), console, priority),
		// 打印在文件中
		zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, priority),
	)
	ch := newCron()
	*loggerMap[key] = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.Hooks(func(entry zapcore.Entry) error {
			select {
			case <-ch:
				if loggerLevel <= entry.Level {
					return l.Rotate()
				}
			default:
			}
			return nil
		}),
	)
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func newCron() chan bool {
	ch := make(chan bool, 1)
	c := cron.New()
	_ = c.AddFunc("0 0 0 * * *", func() {
		select {
		case ch <- true:
		default:
		}
	}) // 秒,分,时,日,月,周 每日零点执行
	c.Start()
	return ch
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func Debug(msg string, fields ...zap.Field) {
	debugLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	infoLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	warnLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	errorLogger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	dpanicLogger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	panicLogger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	fatalLogger.Fatal(msg, fields...)
}
