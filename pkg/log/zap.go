package log

import (
	"fmt"
	"my-ops-admin/global"
	"my-ops-admin/utils"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapCore struct {
	zapcore.Core
}

func Zap() (logger *zap.Logger) {
	level := global.OPS_CONFIG.Zap.Levels()
	zapCores := make([]zapcore.Core, 0, 2)
	if global.OPS_CONFIG.Zap.LogInConsole {
		core := zapcore.NewCore(global.OPS_CONFIG.Zap.ConsoleEncoder(), zapcore.AddSync(os.Stdout), level)
		zapCores = append(zapCores, core)
	}
	if global.OPS_CONFIG.Zap.LogInFile {
		if ok, _ := utils.PathExists(global.OPS_CONFIG.Zap.Director); !ok {
			fmt.Printf("create %v directory\n", global.OPS_CONFIG.Zap.Director)
			err := os.Mkdir(global.OPS_CONFIG.Zap.Director, os.ModePerm)
			if err != nil {
				fmt.Printf("打开日志目录%s失败.\n", global.OPS_CONFIG.Zap.Director)
				panic(err)
			}
		}
		logFile := global.OPS_CONFIG.Zap.Director + "/" + global.OPS_CONFIG.Zap.LogFileName
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("打开日志文件%s失败.\n", logFile)
			panic(err)
		}
		core := zapcore.NewCore(global.OPS_CONFIG.Zap.FileEncoder(), zapcore.AddSync(file), level)
		zapCores = append(zapCores, core)
	}
	logger = zap.New(zapcore.NewTee(zapCores...))
	if global.OPS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
