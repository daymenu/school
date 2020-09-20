package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func mySQL(m *DBConfig) (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.DBName,
			m.Charset,
			m.ParseTime,
			m.Local))
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
