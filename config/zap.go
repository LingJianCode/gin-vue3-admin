package config

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出控制台
	RetentionDay  int    `mapstructure:"retention-day" json:"retention-day" yaml:"retention-day"`    // 日志保留天数
	LogInFile     bool   `mapstructure:"log-in-file" json:"log-in-file" yaml:"log-in-file"`          // 输出文件
	LogFileName   string `mapstructure:"log-file-name" json:"log-file-name" yaml:"log-file-name"`    // 日志文件名称
}

// Levels 根据字符串转化为 zapcore.Levels
func (c *Zap) Levels() zapcore.Level {
	level, err := zapcore.ParseLevel(c.Level)
	if err != nil {
		level = zapcore.DebugLevel
	}
	return level
}

func (c *Zap) ConsoleEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "Time",
		NameKey:       "Name",
		LevelKey:      "Level",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: c.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(c.Prefix + "    " + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // 大写编码器带颜色
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if c.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)

}

func (c *Zap) FileEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		TimeKey:       "Time",
		NameKey:       "Name",
		LevelKey:      "Level",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: c.StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(c.Prefix + "    " + t.Format("2006-01-02 15:04:05.000"))
		},
		EncodeLevel:    zapcore.CapitalLevelEncoder, // 大写编码器
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	if c.Format == "json" {
		return zapcore.NewJSONEncoder(config)
	}
	return zapcore.NewConsoleEncoder(config)

}
