package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {

	// 效验dsn
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/project?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println("MySQL DSN Error:", err)
		return nil, err
	}

	// 效验密码
	if err := db.Ping(); err != nil {
		fmt.Println("MySQL Connection Error:", err)
		return nil, err
	}

	return db, nil;
}

func Select(db *sql.DB) error {
	// 查询数据
	rows, err := db.Query("SELECT `user_id`, `username`, `email`, `status`, `last_login_ip` FROM `admin` LIMIT 1")
	if err != nil {
		fmt.Println("MySQL Query Error:", err)
		return err
	}

	defer rows.Close()

	// 取出数据
	for rows.Next() {
		var (
			UserId      int64
			Username    string
			Email       string
			Status      int
			LastLoginIp string
		)

		// 解析数据
		if err := rows.Scan(&UserId, &Username, &Email, &Status, &LastLoginIp); err != nil {
			return err
		}

		fmt.Println(UserId, Username, Email, Status, LastLoginIp)
	}

	return nil
}

func Query(db *sql.DB, data interface{}) error {
	mapper, err := getInteface(data)
	if err != nil {
		return err
	}

	// 查询数据
	rows, err := db.Query("SELECT * FROM `admin` LIMIT 1")
	if err != nil {
		return err
	}

	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	values := make([]interface{}, len(columns))
	for k, v := range columns {
		if i, ok := mapper[v]; ok {
			values[k] = i;
		}
	}

	for rows.Next() {
		if err := rows.Scan(values...); err != nil {
			return err
		}
	}

	return nil
}

func getInteface(data interface{}) (map[string]interface{}, error) {

	value := reflect.ValueOf(data)

	// 验证必须为指针
	if value.Kind() != reflect.Ptr {
		return nil, errors.New("must pass a pointer, not a value, to StructScan destination")
	}

	if value.IsNil() {
		return nil, errors.New("nil pointer passed to StructScan destination")
	}

	// 必须是一个结构体
	typeElemOf := reflect.TypeOf(data).Elem()
	kind := typeElemOf.Kind()

	if kind != reflect.Struct && kind != reflect.Slice {
		return nil, errors.New("must pass a pointer structure, not a value: " + fmt.Sprintf("%v", typeElemOf.Kind()))
	}

	// 如果是结构体那么
	values := make(map[string]interface{})
	for i := 0; i < typeElemOf.NumField(); i++ {
		column := typeElemOf.Field(i).Tag.Get("db")
		if column != "" {
			values[column] = value.Elem().Field(i).Addr().Interface(); 
		}
	}

	return values, nil
}
