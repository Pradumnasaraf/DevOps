package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	vp := viper.New()

	vp.SetConfigType("json")
	vp.SetConfigName("config")
	// vp.SetConfigType("yaml")
	// vp.SetConfigName("config1")
	vp.AddConfigPath(".")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vp.GetString("user.name"))
	fmt.Println(vp.GetString("user.age"))
}
