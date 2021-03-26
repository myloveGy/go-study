package mysql

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpen(t *testing.T) {
	_ , err := Open()
	assert.NoError(t, err)
}

func TestSelect(t *testing.T) {
	db , err := Open()
	assert.NoError(t, err)

	assert.NoError(t, Select(db))
}

func TestQuery(t *testing.T) {
	db , err := Open()
	assert.NoError(t, err)

	data := &Admin{}
	err = Query(db, data)
	fmt.Println(data, err)

	list := make([]*Admin, 0)

	// 通过反射获取slice元素类型
	model := reflect.New(reflect.ValueOf(list).Type().Elem()).Elem().Interface()
	if v, ok := model.(IModel); ok {
		fmt.Println(v.TableName(), v.PK())
	}

}
