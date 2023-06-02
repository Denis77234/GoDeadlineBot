package main

import (
	"log"
	"katun/internal/database/postgres"
	"katun/internal/bot"

)

func main() {

cfg:= postgres.Config{
	Login: "",
	Password: "",
	Host: "",
	Database: "",
	Ssl: "",
}

katunDB,err:= postgres.New(cfg)
if err!=nil{
	log.Fatal(err)
	return
}

tgbot:= bot.New("YOUR TOKEN",katunDB)

tgbot.Start()


}