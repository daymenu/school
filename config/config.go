package config

import (
	"github.com/daymenu/school/core/db"
)

// AppConfig Config
type AppConfig struct {
	MySQL MySQL
}

// MySQL MySQL
type MySQL struct {
	School db.DBConfig
}
