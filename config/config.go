package config

// Config Config
type Config struct {
	App App
}

// App App
type App struct {
	MySQL MySQL
}

// MySQL MySQL
type MySQL struct {
	Host     string
	Port     int
	Username string
	Password string
}
