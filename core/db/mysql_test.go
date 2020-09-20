package db

import (
	"testing"
)

type User struct {
	Host string
	User string
}

func (User) TableName() string {
	return "user"
}

func TestMySQL(t *testing.T) {
	db, err := New(MySQLType, &DBConfig{
		Host:      "127.0.0.1",
		Port:      3306,
		DBName:    "mysql",
		Username:  "root",
		Password:  "123456",
		Charset:   "utf8",
		ParseTime: true,
		Local:     "Local",
	})
	if err != nil {
		t.Fatal("数据库连接失败:", err)
	}
	user := User{}
	db.First(&user)
	t.Fatal(user)
}
