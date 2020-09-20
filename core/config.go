package core

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var basePath, _ = os.Getwd()

var defaultConfigPath = filepath.Join(basePath, "/resources")

// Config 配置文件
type Config struct {
	config []byte
}

// NewConfig 创建配置类
func NewConfig(v ...string) *Config {
	c := Config{}
	c.Load()
	return &c
}

// Load 加载配置文件
func (c *Config) Load() {
	file, err := os.Open(defaultConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileInfos, err := file.Readdir(0)
	for _, f := range fileInfos {
		if f.IsDir() {
			continue
		}
		fpath := filepath.Join(defaultConfigPath, f.Name())
		cbyte, err := ioutil.ReadFile(fpath)
		if err != nil {
			continue
		}
		c.config = cbyte
	}
}

func (c *Config) GetConfig(v interface{}) {
	yaml.Unmarshal(c.config, v)
}
