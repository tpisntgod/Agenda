package mylog

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var logDivPath = "src/github.com/tpisntgod/Agenda/Json/Log/"
var logFilePath = time.Now().Format("2006-01-02") + ".txt"

func init() {
	logDivPath = filepath.Join(os.Getenv("GOPATH"), logDivPath)
}

func getFileHandle() *os.File {
	if _, err := os.Open(logDivPath + logFilePath); err != nil {
		os.Create(logDivPath + logFilePath)
	}

	// 以追加模式打开文件,并向文件写入
	fi, _ := os.OpenFile(logDivPath+logFilePath, os.O_RDWR|os.O_APPEND, 0)
	return fi
}

// AddLog : 添加记录
func AddLog(user string, command string, oldStr string, newStr string) {
	file := getFileHandle()
	l := log.New(file, "[INFO]", log.Ltime)
	outStr := ""
	if user != "" {
		outStr += "User:" + user + "  "
	}
	if command != "" {
		outStr += "Command:" + command + "\n"
	}
	if oldStr != "" {
		outStr += "From:" + oldStr + "\n"
	}
	if newStr != "" {
		outStr += "To:" + newStr + "\n"
	}
	l.Print(outStr)
	file.Close()
}
