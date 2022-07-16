package main

import (
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("your token")
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			photoBytes, _ := ioutil.ReadFile(fmt.Sprintf("./img/%v", choice_random_gif()))

			photoFileBytes := tgbotapi.FileBytes{
				Name:  "picture",
				Bytes: photoBytes,
			}
			chatID := update.Message.Chat.ID
			bot.Send(tgbotapi.NewPhoto(int64(chatID), photoFileBytes))

			write_id(chatID)

		}
	}

}

func choice_random_gif() string {

	lst := []string{}

	files, _ := ioutil.ReadDir("./img")

	for _, file := range files {

		lst = append(lst, file.Name())

	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(lst))
	pick := lst[randomIndex]

	return pick

}

func write_id(get_id_user int64) {

	database, _ := sql.Open("sqlite3", "./test.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS id_tg_user (id INTEGER PRIMARY KEY, id_user INTEGER UNIQUE)")
	statement.Exec()

	statement, _ = database.Prepare("INSERT INTO id_tg_user (id_user) VALUES (?)")
	statement.Exec(get_id_user)

}
