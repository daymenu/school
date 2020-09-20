package model

import "github.com/jinzhu/gorm"

// SchoolInfo 学校信息
type SchoolInfo struct {
	gorm.Model
	Name string
}
