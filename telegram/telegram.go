package telegram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func SendMessage(chat_id string, text string) {

	token := ""

	telegramUrl, err := url.Parse("https://api.telegram.com/bot" + token + "/sendMessage")

	if err != nil {
		log.Println(err)
	}

	values := telegramUrl.Query()
	values.Add("chat_id", chat_id)
	values.Add("text", chat_id)
	values.Add("parse_mode", chat_id)

	response, err := http.Get(telegramUrl.String())
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

}
