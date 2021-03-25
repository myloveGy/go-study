package mysql

import "time"

type IModel interface {
	TableName() string
	PK() string
}

type Admin struct {
	UserId        int64     `db:"user_id"`
	Username      string    `db:"username"`
	Password      string    `db:"password"`
	Email         string    `db:"email"`
	Status        int       `db:"status"`
	LastLoginIp   string    `db:"last_login_ip"`
	LastLoginTime time.Time `db:"last_login_time"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (*Admin) TableName() string {
	return "admin";
}

func (*Admin) PK() string {
	return "user_id"
}
