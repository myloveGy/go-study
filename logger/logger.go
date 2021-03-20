package logger

import (
	"encoding/json"
	"io"
	"runtime"
	"strings"
	"sync"
	"time"
)

type LogLevel uint

type logContent struct {
	Time     string      `json:"time"`    // 时间
	FileName string      `json:"file"`    // 文件名称
	Line     int         `json:"line"`    // 行号
	FuncName string      `json:"func"`    // 函数名称
	Level    string      `json:"level"`   // 日志级别
	Message  string      `json:"message"` // 消息
	Content  interface{} `json:"content"` // 内容信息
}

const (
	Debug   = "DEBUG"
	Notice  = "NOTICE"
	Trace   = "Trace"
	Info    = "INFO"
	Warning = "WARNING"
	Error   = "ERROR"
	Fatal   = "FATAL"

	LOG_LEVEL_DEBUG LogLevel = iota
	LOG_LEVEL_NOTICE
	LOG_LEVEL_TRACE
	LOG_LEVEL_INFO
	LOG_LEVEL_WARN
	LOG_LEVEL_ERROR
	LOG_LEVEL_FATAL

	// 默认输出格式
	DefaultTemplate = "[{time}] {level}: {message} {content}"

	// 默认的时间格式
	DefaultFormat = "2006-01-02 15:04:05"
)

func parseLogLevel(level string) LogLevel {
	switch level {
	case "debug", Debug, "Debug":
		return LOG_LEVEL_DEBUG
	case "notice", Notice, "Notice":
		return LOG_LEVEL_NOTICE
	case "info", Info, "Info":
		return LOG_LEVEL_INFO
	case "trace", "TRACE", "Trace":
		return LOG_LEVEL_TRACE
	case "warning", "WARNING", "Warning":
		return LOG_LEVEL_WARN
	case "error", "ERROR", "Error":
		return LOG_LEVEL_ERROR
	case "fatal", "FATAL", "Fatal":
		return LOG_LEVEL_FATAL
	default:
		return LOG_LEVEL_DEBUG
	}
}

type Logger struct {
	write    io.Writer  // 写类型
	format   string     // 日志时间格式
	template string     // 日志输出格式
	level    LogLevel   // 日志级别
	json     bool       // 是否json格式记录日志
	mutex    sync.Mutex // 锁
}

func NewLogger(write io.Writer, level string) *Logger {
	return &Logger{
		write:    write,
		format:   DefaultFormat,
		template: DefaultTemplate,
		level:    parseLogLevel(level),
	}
}

func (l *Logger) SetLogLevel(info string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.level = parseLogLevel(info)
}

func (l *Logger) SetJson(json bool) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.json = json
}

func (l *Logger) Debug(message string, content interface{}) {
	l.log(Debug, message, content)
}

func (l *Logger) Notice(message string, content interface{}) {
	l.log(Notice, message, content)
}

func (l *Logger) Info(message string, content interface{}) {
	l.log(Info, message, content)
}

func (l *Logger) Warning(message string, content interface{}) {
	l.log(Warning, message, content)
}

func (l *Logger) Error(message string, content interface{}) {
	l.log(Error, message, content)
}

func (l *Logger) Trace(message string, content interface{}) {
	l.log(Trace, message, content)
}

func (l *Logger) Fatal(message string, content interface{}) {
	l.log(Fatal, message, content)
}

func (l *Logger) log(level, message string, content interface{}) {

	// 设置了级别
	if parseLogLevel(level) < l.level {
		return
	}

	nowTime := time.Now().Format(l.format)

	var writeString string
	if l.json {
		funcName, fileName, line := getCaller(3)

		jsonData, _ := json.Marshal(&logContent{
			Time:     nowTime,
			FuncName: funcName,
			Line:     line,
			FileName: fileName,
			Level:    level,
			Message:  message,
			Content:  content,
		})

		writeString = string(jsonData)
	} else {
		jsonData, _ := json.Marshal(content)
		// 替换时间、等级、消息、内容
		writeString = strings.Replace(l.template, "{time}", nowTime, -1)
		writeString = strings.Replace(writeString, "{level}", level, -1)
		writeString = strings.Replace(writeString, "{message}", message, -1)
		writeString = strings.Replace(writeString, "{content}", string(jsonData), -1)
	}

	_, _ = l.write.Write([]byte(writeString + "\n"))
}

func getCaller(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", "", 0
	}

	return runtime.FuncForPC(pc).Name(), file, line
}
