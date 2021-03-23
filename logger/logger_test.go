package logger

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger(os.Stdout, "", 100)
	assert.Equal(t, os.Stdout, logger.write)
	assert.Equal(t, DefaultTemplate, logger.template)
}

func TestInfo(t *testing.T) {
	logger := NewLogger(os.Stdout, "info", 100)
	logger.SetJson(true)
	logger.SetFormat("2006/01/02 15:04:05")
	logger.Debug("我的debug", "测试内容")
	logger.Info("我的测试", nil)
	logger.Trace("我的测试", map[string]interface{}{
		"username": "jinxing.liu",
		"age":      18,
	})
	logger.Warning("我的warning", "测试内容呢")
	logger.Error("出现错误了", ":)")
	logger.Fatal("什么信息", 123)
	fmt.Println("等待执行完成", logger.Close(), logger.Wait())

	file, err := os.OpenFile(fmt.Sprintf("./%s.log", time.Now().Format("20060102")), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	assert.NoError(t, err)
	defer file.Close()

	logger2 := NewLogger(file, "", 100)
	logger2.SetTemplate("[{time}] {level}: {message} {content}")
	logger2.Notice("测试Notcie", nil)
	logger2.Warning("测试Warning", "就是测试的数据")
	logger2.Info("测试Info", 123232323)
	logger2.Debug("测试Debug", "不知道")
	logger2.Error("测试Error", "不知道")
	logger2.Fatal("测试Fatal", "不知道")
	fmt.Println("等待执行完成", logger2.Close(), logger2.Wait())
}

func Test_getCaller(t *testing.T) {
	tests := []struct {
		name string
		arg  int
	}{
		{name: "测试", arg: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			funcName, fileName, line := getCaller(tt.arg)
			assert.NotEmpty(t, funcName)
			assert.NotEmpty(t, fileName)
			assert.NotEmpty(t, line)
			fmt.Println(funcName, fileName, line)
		})
	}
}

func Test_parseLogLevel(t *testing.T) {
	tests := []struct {
		name string
		want LogLevel
	}{

		{name: "info", want: LOG_LEVEL_INFO},
		{name: "Info", want: LOG_LEVEL_INFO},
		{name: "INFO", want: LOG_LEVEL_INFO},
		{name: "Warning", want: LOG_LEVEL_WARN},
		{name: "warning", want: LOG_LEVEL_WARN},
		{name: "WARNING", want: LOG_LEVEL_WARN},
		{name: "Trace", want: LOG_LEVEL_TRACE},
		{name: "trace", want: LOG_LEVEL_TRACE},
		{name: "TRACE", want: LOG_LEVEL_TRACE},
		{name: "debug", want: LOG_LEVEL_DEBUG},
		{name: "error", want: LOG_LEVEL_ERROR},
		{name: "fatal", want: LOG_LEVEL_FATAL},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLogLevel(tt.name); got != tt.want {
				t.Errorf("parseLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
