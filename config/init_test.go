package config

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	config, err := Load("./config.ini")
	assert.NoError(t, err)
	assert.Equal(t, "jinxing.liu", config.data["app_name"])
}

func TestParse(t *testing.T) {
	config, err := Load("./config.ini")
	assert.NoError(t, err)

	// 解析app
	appConfig := &Config{}
	err = config.Parse(appConfig)
	assert.NoError(t, err)
	fmt.Printf(
		"app_name = %#v\ndebug = %#v\nmysqlConfig = %#v\nredisConfig = %#v\n",
		appConfig.AppName,
		appConfig.Debug,
		appConfig.MySqlConfig,
		appConfig.RedisConfig,
	)

	// 外层赋值判断
	assert.Equal(t, "jinxing.liu", appConfig.AppName)
	assert.Equal(t, true, appConfig.Debug)

	// 内部引用组合
	assert.Equal(t, "root", appConfig.MySqlConfig.Username)
	assert.Equal(t, "root:123456@tcp(127.0.0.1:3306)/project?charsetutf8&parseTimeTrue&locAsia%2FShanghai", appConfig.MySqlConfig.Dsn)
	assert.Equal(t, 6379, appConfig.RedisConfig.Port)
}

func TestParseStruct(t *testing.T) {
	config, err := Load("./config.ini")
	assert.NoError(t, err)

	arr := &struct {
		MySqlConfig `ini:"mysql"`
		RedisConfig `ini:"redis"`
	}{}

	assert.NoError(t, config.Parse(arr))
	fmt.Printf("mysql = %#v\nredis = %#v\n", arr.MySqlConfig, arr.RedisConfig)
}

func TestParseSingle(t *testing.T) {
	config, err := Load("./config.ini")
	assert.NoError(t, err)

	// 解析错误
	assert.Error(t, config.Parse("123456"))
	var n = 123
	assert.Error(t, config.Parse(&n))

	// 单独解析mysql
	mysqlConfig := &MySqlConfig{}
	assert.NoError(t, config.Parse(mysqlConfig, "mysql"))
	fmt.Printf("mysql = %#v\n", mysqlConfig)
	assert.Equal(t, "127.0.0.1", mysqlConfig.Host)
	assert.Equal(t, "123456", mysqlConfig.Password)

	// 单独解析redis
	redisConfig := &RedisConfig{}
	assert.NoError(t, config.Parse(redisConfig, "redis"))
	fmt.Printf("mysql = %#v\n", redisConfig)
	assert.Equal(t, 6379, redisConfig.Port)
}

func TestGetValue(t *testing.T) {
	config, err := Load("./config.ini")
	assert.NoError(t, err)

	// 获取map
	mysql, has := config.GetValue("mysql")
	assert.Equal(t, true, has)
	fmt.Println(mysql)

	// .语法直接获取二级
	mysqlHost, has := config.GetValue("mysql.host")
	assert.Equal(t, true, has)
	assert.Equal(t, "127.0.0.1", mysqlHost)

	appName, has := config.GetValue("app_name")
	assert.Equal(t, true, has)
	assert.Equal(t, "jinxing.liu", appName)

	_, has = config.GetValue("hahaha")
	assert.Equal(t, false, has)

	nilValue, has := config.GetValue("")
	assert.Equal(t, false, has)
	assert.Nil(t, nilValue)
}

func TestGet(t *testing.T) {
	config, err := Load("./config.ini")
	assert.NoError(t, err)

	// 获取map
	mysql, has := config.Get("mysql", "")
	assert.Equal(t, true, has)
	fmt.Println(mysql)

	// .语法直接获取二级
	mysqlHost, has := config.Get("mysql.host", "")
	assert.Equal(t, true, has)
	assert.Equal(t, "127.0.0.1", mysqlHost)

	appName, has := config.Get("app_name", "")
	assert.Equal(t, true, has)
	assert.Equal(t, "jinxing.liu", appName)

	defaultValue, has := config.Get("hahaha", "456")
	assert.Equal(t, false, has)
	assert.Equal(t, "456", defaultValue)

	value, has := config.Get("", "")
	assert.Equal(t, false, has)
	assert.Equal(t, "", value)
}

func TestSet(t *testing.T) {
	config := &IniConfig{data: map[string]interface{}{
		"name": 456,
		"info": map[string]string{"username": "jinxing.liu"},
	}}

	assert.NoError(t, config.Set("username", "jinxing.liu"))
	assert.NoError(t, config.Set("age", "25"))
	assert.NoError(t, config.Set("info.test", "123"))

	username, _ := config.Get("username", "")
	assert.Equal(t, "jinxing.liu", username)

	age, _ := config.Get("age", "")
	assert.Equal(t, "25", age)

	assert.Error(t, config.Set("name.user", "456"))
	assert.Error(t, config.Set("", "789"))
	assert.Error(t, config.Set("haha.username", "789"))
}

func TestRegexp(t *testing.T) {
	fmt.Println(regexp.MustCompile(`\$\{(\w{0,}\.{0,}\w{0,})\}`).FindAllString("${username}12121212${host}", -1))
}
