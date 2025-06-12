package config

import (
	"errors"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
	"path"
	"runtime"
)

type Config struct {
	Database Database `yaml:"DATABASE"`
}

type Database struct {
	User     string `yaml:"USER"`
	Password string `yaml:"PASSWORD"`
	Name     string `yaml:"NAME"`
	Host     string `yaml:"HOST"`
	Port     string `yaml:"PORT"`
}

func readFile(cfg *Config) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		processError(errors.New("unable to get config file directory"))
	}

	f, err := os.Open(path.Join(path.Dir(filename), "/config.yaml"))
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func Get() *Config {
	var cfg Config
	readFile(&cfg)
	readEnv(&cfg)
	return &cfg
}

func processError(err error) {
	panic(err)
}
