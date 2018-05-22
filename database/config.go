package database

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

var lookupPaths = []string{"", "./config", "/config", "../", "../config", "../..", "../../config"}

// ConfigName is the name of the YAML databases config file
var ConfigName = "database.yml"

type Config struct {
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Pool     int    `yaml:"pool"`
	Dialect  string `yaml:"dialect"`
}

func (db Config) toString() string {
	switch db.Dialect {
	case "mysql":
		return fmt.Sprintf(
			"%v:%v@%v/%v?charset=utf8&parseTime=True",
			db.User,
			db.Password,
			db.Host,
			db.Database,
		)
	case "sqlite3":
		return db.Database
	}
	return ""
}

func findConfigPath() (string, error) {
	for _, p := range lookupPaths {
		path, _ := filepath.Abs(filepath.Join(p, ConfigName))
		if _, err := os.Stat(path); err == nil {
			return path, err
		}
	}
	return "", errors.New("[Burrow]: Tried to load configuration file, but couldn't find it")
}

func loadFrom(env string) Config {
	path, err := findConfigPath()
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]Config)
	yamlFile, _ := ioutil.ReadFile(path)
	yaml.Unmarshal(yamlFile, &m)

	return m[env]
}
