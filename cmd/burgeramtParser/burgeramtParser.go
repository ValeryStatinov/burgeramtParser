package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ValeryStatinov/burgeramtParser/internal/parser"
	"github.com/ValeryStatinov/burgeramtParser/internal/tg"
)

func main() {
	tgToken := os.Getenv("TG_TOKEN")
	chatId := os.Getenv("CHAT_ID")

	dls := parser.NewDrivingLicenceSpider()
	tgClient := tg.NewTgClient(tgToken)

	ticker := time.NewTicker(time.Minute * 3)

	tick := func() {
		fmt.Println("start crawling...")
		dates, err := dls.Crawl()
		if err != nil {
			tgClient.SendMessage(err.Error(), chatId)

			return
		}

		for _, d := range dates {
			tgClient.SendMessage(d, chatId)
		}
	}

	err := tgClient.SendMessage("starting crawling bot", chatId)
	if err != nil {
		log.Fatal("failed to send message to telegram bot")
		return
	}

	tick()
	for range ticker.C {
		tick()
	}
}
