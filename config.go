// Copyright (c) 2018 Chun-Koo Park

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %s\n", err))
	}

	dbName := viper.GetString("database.name")
	dbUser := viper.GetString("database.user")
	//dbPassword := viper.GetString("database.password")

	if len(dbName) == 0 || len(dbUser) == 0 {
		panic("Please enter database name/user/password in config file.")
	}

}
