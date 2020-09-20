package core

import (
	"testing"

	"github.com/daymenu/school/config"
)

func TestNew(t *testing.T) {
	appConfig := config.AppConfig{}
	c := NewConfig()
	c.GetConfig(&appConfig)
	t.Fatal(appConfig)
}
