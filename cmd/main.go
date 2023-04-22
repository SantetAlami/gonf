package main

import (
	"fmt"
	"log"

	l "github.com/SantetAlami/go-conf-loader"
)

type Config struct {
	Db struct {
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"db"`
}

func main() {
	conf := Config{}
	if err := l.Load(&conf, "cmd/configs/config.yaml", ""); err != nil {
		log.Panic(err)
	}
	fmt.Println(conf.Db.User)
}