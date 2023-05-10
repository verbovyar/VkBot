package main

import (
	"bot/Config"
	"bot/Repository/DataBase"
	"bot/Service"
	"log"
)

func main() {
	conf, err := Config.LoadConfig("././Config")
	if err != nil {
		log.Fatalf("%v", err)
	}

	Repo := DataBase.New()
	Service.New(Repo)

	log.Println("start cmd")
	bot, err := Service.Init(conf.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	Service.AddHandlers(bot)

	if err := bot.Run(); err != nil {
		log.Panic(err)
	}
}
