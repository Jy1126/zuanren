package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Conf struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

var (
	Config *Conf
)

func GetConf() {
	dir, err := os.Getwd()
	fmt.Println(filepath.Join(dir, "", "conf.yaml"))
	yamlFile, err := ioutil.ReadFile(filepath.Join(dir, "", "conf.yaml"))
	if err != nil {
		log.Printf("conf.yaml get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}

	// 环境变量覆盖
	dbHost := os.Getenv("DB_Host")
	dbPort := os.Getenv("DB_Port")
	dbUser := os.Getenv("DB_User")
	dbPassword := os.Getenv("DB_Password")
	if dbHost != "" {
		Config.Host = dbHost

	}
	if dbPort != "" {
		Config.Port = dbPort

	}
	if dbUser != "" {
		Config.User = dbUser

	}
	if dbPassword != "" {
		Config.Password = dbPassword

	}

}
