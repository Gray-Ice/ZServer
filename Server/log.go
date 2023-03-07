package Server

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"runtime"
	"strings"
)

var Logger *logrus.Logger

type hook struct{}

func (h *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (h *hook) Fire(e *logrus.Entry) error {
	ljLogger := &lumberjack.Logger{
		Filename:   getLogPath(),
		MaxSize:    10,
		MaxBackups: 2,
	}
	writer := io.MultiWriter(os.Stdout, ljLogger)
	e.Logger.SetOutput(writer)
	return nil
}

func getLogPath() string {
	slash := ""
	if runtime.GOOS == "windows" {
		slash = "\\"
	} else {
		slash = "/"
	}

	path, err := os.Executable()
	if err != nil {
		return "." + slash
	}

	// 跨平台, windows采用\,其他平台采用/
	pos := strings.LastIndex(path, slash)
	if pos <= 0 {
		return "." + slash
	}
	return path[:pos] + slash + "zlog.log"
}

// 初始化Logger
func init() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)                 // 显示调用文件
	Logger.SetFormatter(&logrus.JSONFormatter{}) // 以JSON格式显示日志
	Logger.AddHook(&hook{})
}
