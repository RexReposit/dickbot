package main

import (
	"bot/src/handlers"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	dataHandler := handlers.Data{}
	dataHandler.InitDB()

	pref := tele.Settings{
		Token:  "7757765456:AAFpFXhbi9XCfgRt7P3OT3F_jrBBplubWZA",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/dick", dataHandler.Dick)
	b.Handle("/top_dick", dataHandler.TopDick)
	b.Handle("/clear_stats", dataHandler.ClearStatistics)

	b.Start()
}
