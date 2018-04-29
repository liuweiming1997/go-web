package db

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type dbConfig struct {
	DB_Host     string `default:"39.108.100.86"`
	DB_Port     string `default:"3306"`
	DB_User     string `default:"root"`
	DB_Password string `default:"dc"`
	DB_Name     string `default:"homework"`
}

func getDBConfig() (*dbConfig, error) {
	var conf dbConfig
	err := envconfig.Process("DC", &conf)
	fmt.Println(conf)
	return &conf, err
}
