package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"meal-planner/app"
)

func main() {
	viper.SetConfigFile("configs/main.yml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	app.Run()
}
