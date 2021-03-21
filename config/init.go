package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type IniConfig struct {
	data  map[string]interface{} // 配置内容
	mutex sync.Mutex             // 锁
}

func Load(filename string) (*IniConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// 读取配置信息
	config := &IniConfig{
		data: make(map[string]interface{}),
	}

	reader := bufio.NewReader(file)

	// 分组名称
	groupName := ""
	line := 0

	for {
		line += 1
		lineString, err := reader.ReadString('\n')
		lineString = strings.TrimSpace(lineString)
		if lineString == "" && err == nil {
			continue
		}

		// 开始解析数据
		if lineString[0] == '[' {
			// [ 开头但是不是]结尾，抛出错误
			if lineString[len(lineString)-1] != ']' {
				return nil, fmt.Errorf("config group name read error line: %d", line)
			}

			// 获取内部内容为groupName
			groupName = string(lineString[1 : len(lineString)-1])
			if err := config.SetMap(groupName, make(map[string]string)); err != nil {
				return nil, fmt.Errorf("config group name set error %s line: %d", err.Error(), line)
			}
		} else if strings.Contains(lineString, "=") {
			keyNames := strings.Split(lineString, "=")
			if len(keyNames) != 2 || strings.TrimSpace(keyNames[0]) == "" {
				return nil, fmt.Errorf("config value read error line: %d", line)
			}

			name := strings.TrimSpace(keyNames[0])
			value := strings.TrimSpace(keyNames[1])
			if groupName != "" {
				name = fmt.Sprintf("%s.%s", groupName, name)
			}

			if err := config.Set(name, value); err != nil {
				return nil, fmt.Errorf("config value set error %s line: %d", err.Error(), line)
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			// 存在错误退出
			return nil, err
		}
	}

	return config, nil
}

func (i *IniConfig) Parse(data interface{}, names ...string) error {
	// 验证传递的数据
	dataType := reflect.TypeOf(data)

	// 1. 必须是一个指针类型
	if dataType.Kind() != reflect.Ptr {
		return errors.New("data not pointer")
	}

	// 2. 必须是一个结构体类型
	nameElem := dataType.Elem()
	if nameElem.Kind() != reflect.Struct {
		return errors.New("data not struct")
	}

	dataValue := reflect.ValueOf(data).Elem()

	for num, length := 0, dataValue.NumField(); num < length; num++ {
		name := nameElem.Field(num).Tag.Get("ini")
		if name == "" {
			continue
		}

		if len(names) == 1 {
			name = fmt.Sprintf("%s.%s", names[0], name)
		}

		fieldValue := dataValue.Field(num)  // value 字段value 类型
		fieldValueKind := fieldValue.Kind() // 字段类型

		// 不能设置值
		if !fieldValue.CanSet() {
			continue
		}

		// key 存在map中，并且能够赋值
		if value, has := i.GetValue(name); has {
			if strValue, ok := value.(string); ok {
				if err := i.setReflectValue(fieldValueKind, fieldValue, strValue); err != nil {
					return err
				}

				continue
			}

			// 是map[string]string 类型，并且是指针类型或者结构体类型
			if mapValue, ok := value.(map[string]string); ok {
				// 指针类型
				if fieldValueKind == reflect.Ptr {
					// 不是结构体类型，不处理
					if fieldValue.Type().Elem().Kind() != reflect.Struct {
						continue
					}

					// 空值的情况nil, 需要创建一个
					if fieldValue.IsNil() {
						fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
					}

					if err := i.setStructReflectValue(fieldValue.Elem(), mapValue); err != nil {
						return err
					}

					continue
				}

				// 结构体类型
				if fieldValueKind == reflect.Struct {
					if err := i.setStructReflectValue(fieldValue, mapValue); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (i *IniConfig) GetValue(name string) (interface{}, bool) {
	// 传递的name为空
	if name == "" {
		return nil, false
	}

	i.mutex.Lock()
	defer i.mutex.Unlock()

	keys := strings.Split(name, ".")

	// 不存在第一个值，那么直接返回默认值
	value, has := i.data[keys[0]]
	if !has {
		return nil, false
	}

	// 如果是 name 那么直接返回
	if len(keys) == 1 {
		return value, has
	}

	// mysql.host
	if mapValue, ok := value.(map[string]string); ok {
		tempValue, has := mapValue[keys[1]]
		return tempValue, has
	}

	return nil, false
}

func (i *IniConfig) Get(name, defaultValue string) (string, bool) {
	// 存在并且是字符串类型，那么直接返回
	if value, has := i.GetValue(name); has {
		return fmt.Sprintf("%v", value), true
	}

	// 不存在直接返回默认值
	return defaultValue, false
}

func (i *IniConfig) Set(name, value string) error {
	if name == "" {
		return errors.New("name is empty")
	}

	i.mutex.Lock()
	defer i.mutex.Unlock()

	keyNames := strings.Split(name, ".")
	if len(keyNames) == 1 {
		i.data[keyNames[0]] = value
		return nil
	}

	if _, ok := i.data[keyNames[0]]; !ok {
		return fmt.Errorf("name: %s not set value", keyNames[0])
	}

	data := i.data[keyNames[0]]
	v, ok := data.(map[string]string)
	if !ok {
		return fmt.Errorf("name: %s not is map[string]string %T", keyNames[0], i.data[keyNames[0]])
	}

	v[keyNames[1]] = value
	i.data[keyNames[0]] = v

	return nil
}

func (i *IniConfig) SetMap(name string, value map[string]string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if _, ok := i.data[name]; ok {
		return fmt.Errorf("%s exists", name)
	}

	i.data[name] = value
	return nil
}

func (i *IniConfig) setReflectValue(reflectKind reflect.Kind, reflectValue reflect.Value, strValue string) error {
	switch reflectKind {
	case reflect.String:
		reflectValue.SetString(strValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return err
		}

		reflectValue.SetInt(intValue)
	case reflect.Float32, reflect.Float64:
		floatValue, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return err
		}

		reflectValue.SetFloat(floatValue)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(strValue)
		if err != nil {
			return err
		}

		reflectValue.SetBool(boolValue)
	}

	return nil
}

func (i *IniConfig) setStructReflectValue(reflectValue reflect.Value, mapValue map[string]string) error {
	reflectType := reflectValue.Type()
	for num, length := 0, reflectValue.NumField(); num < length; num++ {
		name := reflectType.Field(num).Tag.Get("ini")
		if name == "" {
			continue
		}

		strValue, ok := mapValue[name]
		if !ok {
			continue
		}

		structValue := reflectValue.Field(num)
		if err := i.setReflectValue(structValue.Type().Kind(), structValue, strValue); err != nil {
			return err
		}
	}

	return nil
}
