package util

import (
	"reflect"
	"strconv"
	"time"
)

func SetStructFiledValue(reflectValue reflect.Value, strValue string) error {
	switch reflectValue.Type().Kind() {
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
	case reflect.Struct:
		switch reflectValue.Interface().(type) {
		case time.Time:
			t, err := time.Parse("2006-01-02 15:04:05", strValue)
			if err != nil {
				return err
			}

			reflectValue.Set(reflect.ValueOf(t))
		}
	}

	return nil
}
