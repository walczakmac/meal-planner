package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"meal-planner/internal/app"
)

func init() {
	viper.SetConfigFile("configs/main.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.Run()
}
