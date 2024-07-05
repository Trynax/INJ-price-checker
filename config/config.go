package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token string `json:Token"`
	BotPrefix string `json:"BotPrefix"`
}

func ReadConfig () error{
	fmt.Println("Reading the config file.....")
	file,err := ioutil.ReadFile("./config.json")
	if err != nil{
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	err= json.Unmarshal(file, &config)
	if err != nil{
		fmt.Println(err.Error())
		return err
	}

	Token = "MTI1NzMxOTgzNzQxMjEwMjE1NA.GY2EaZ.zmzH-SM8UptTJpjPNDfF7TMl_0Mep1nTjN1tCY"
	BotPrefix = config.BotPrefix

	return nil
}  