package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type dbtype int

const (
	MySQLType dbtype = iota
)

type DBConfig struct {
	Host      string
	Port      int
	DBName    string
	Username  string
	Password  string
	Charset   string
	ParseTime bool
	Local     string
}

func New(t dbtype, config *DBConfig) (db *gorm.DB, err error) {
	switch t {
	case MySQLType:
		db, err = mySQL(config)
	default:
		db, err = nil, errors.New("不支持该数据库驱动方式")
	}

	return db, err
}
