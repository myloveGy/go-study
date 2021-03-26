package redis

import (
	"errors"
	"reflect"
	"study/util"
)

func MapStringToStruct(m map[string]string, data interface{}, args ...string) error {

	// 判断必须要是结构体对象
	valueOf := reflect.ValueOf(data)
	if valueOf.Kind() != reflect.Ptr {
		return errors.New("must pointer")
	}

	elem := valueOf.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("must struct")
	}

	typeElem := valueOf.Type().Elem()
	for i := 0; i < elem.NumField(); i++ {
		keyName := util.Snake(typeElem.Field(i).Name)
		strValue, ok := m[keyName]
		if !ok {
			continue
		}

		// 执行赋值
		if err := util.SetStructFiledValue(elem.Field(i), strValue); err != nil {
			return err
		}
	}

	return nil
}
