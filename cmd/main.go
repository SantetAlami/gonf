package main

import (
	"fmt"
	"log"

	"github.com/SantetAlami/gonf"
)

type Config struct {
	Db struct {
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"db"`
}

func main() {
	conf := Config{}
	if err := gonf.Load(&conf, "cmd/configs/config.yaml", ""); err != nil {
		log.Panic(err)
	}
	fmt.Println(conf.Db.User)
}