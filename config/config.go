package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Throttle struct {
	Rate      int           `yaml:"rate"`
	Burst     int           `yaml:"burst"`
	ExpiresIn time.Duration `yaml:"expires_in"`
}

type GZip struct {
	Level int `yaml:"level"`
}

type Server struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	Throttle Throttle `yaml:"throttle"`
	GZip     GZip     `yaml:"gzip"`
}

type Config struct {
	Server    Server `yaml:"server"`
	Debug     bool   `yaml:"debug"`
	PublicDir string `yaml:"public_dir"`
	DataDir   string `yaml:"data"`
}

func validateConfig(conf *Config) {
	data_dir_stat, err := os.Stat(conf.DataDir)
	if err != nil {
		log.Fatal(err)
	}
	if !data_dir_stat.IsDir() {
		log.Fatal("\"data\" path is not a directory")
	}
	if conf.PublicDir != "" {
		public_dir_stat, err := os.Stat(conf.PublicDir)
		if err != nil {
			log.Fatal(err)
		}
		if !public_dir_stat.IsDir() {
			log.Fatal("\"public_dir\" path is not a directory")
		}
	}
}

func ReadConfig(path string) *Config {
	f, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var conf Config
	err = yaml.Unmarshal(f, &conf)
	if err != nil {
		log.Fatal(err)
	}
	validateConfig(&conf)
	return &conf
}
