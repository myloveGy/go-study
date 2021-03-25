package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Open() error {

	// 效验dsn
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/project?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println("MySQL DSN Error:", err)
		return err
	}

	defer db.Close()

	// 效验密码
	if err := db.Ping(); err != nil {
		fmt.Println("MySQL Connection Error:", err)
		return err
	}

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

		if err := rows.Scan(&UserId, &Username, &Email, &Status, &LastLoginIp); err != nil {
			return err
		}

		fmt.Println(UserId, Username, Email, Status, LastLoginIp)
	}

	return nil
}
