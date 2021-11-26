package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DbConfig struct {
	Addr     string
	Port     int
	Username string
	Name     string
	Password string
	MaxCon   int
	MinCon   int
}

func GetConfig() DbConfig {
	config := DbConfig{}
	file := "./config.json"
	data, err := ioutil.ReadFile(file)
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}

func GetDSN() string {
	config := GetConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Addr,
		config.Port,
		config.Username,
		config.Name,
		config.Password,
	)

	return dsn
}

func Max() int {
	config := GetConfig()

	return config.MaxCon
}

func Min() int {
	config := GetConfig()

	return config.MinCon
}
