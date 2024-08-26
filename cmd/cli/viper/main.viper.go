package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fail to load configuration %w \n", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable")
	}

	fmt.Print("Port:", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database with User: %s, Password: %s", db.User, db.Password)
	}
}