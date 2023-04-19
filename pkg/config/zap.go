package config

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go-template/global"
	"go-template/pkg/file"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"strings"
	"time"
)

type _zap struct {
}

var Zap = new(_zap)

func (z *_zap) InitZap() {
	var director = viper.GetString("zap.director")
	if ok, _ := file.PathExists(director); !ok {
		fmt.Printf("create %v directory\n", director)
		_ = os.Mkdir(director, os.ModePerm)
	}
	cores := z.GetZapCores()
	global.Log = zap.New(zapcore.NewTee(cores...))

	if viper.GetBool("zap.show-line") {
		global.Log = global.Log.WithOptions(zap.AddCaller())
	}
}

func (z *_zap) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(viper.GetString("zap.director"), "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(viper.GetDuration("zap.max-age"))*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if viper.GetBool("zap.log-in-console") {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}

func (z *_zap) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "times",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  viper.GetString("zap.stacktrace-key"),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.ZapEncodeLevel(viper.GetString("zap.encode-level")),
		EncodeTime:     z.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

func (z *_zap) GetEncoder() zapcore.Encoder {
	if viper.GetString("zap.format") == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
}

func (z *_zap) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := z.GetWriteSyncer(l.String())
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}

	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

func (z *_zap) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.TransportLevel(viper.GetString("zap.level")); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

func (z *_zap) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(viper.GetString("zap.prefix") + t.Format("2006/01/02 - 15:04:05.000"))
}

func (z *_zap) ZapEncodeLevel(encoderLevel string) zapcore.LevelEncoder {
	switch {
	case encoderLevel == "LowercaseLevelEncoder":
		// 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case encoderLevel == "LowercaseColorLevelEncoder":
		// 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case encoderLevel == "CapitalLevelEncoder":
		// 大写编码器
		return zapcore.CapitalLevelEncoder
	case encoderLevel == "CapitalColorLevelEncoder":
		// 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

func (z *_zap) TransportLevel(level string) zapcore.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

type Output struct {
	Logger *zap.Logger
}

func (t *Output) Write(p []byte) (n int, err error) {
	t.Logger.Info(string(p))
	return len(p), nil
}
