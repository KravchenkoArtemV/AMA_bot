package telegramClient

import (
	weather "AMA_bot/pkg/weatherAPI"
	"fmt"
	"log"
	"time"
)

// хост является константой
const Host string = "api.telegram.org"

func TgClient(Token string) {
	var token = Token
	// создаем клиента для обработки ответа от тг и отправки структуры на сервер
	client := NewClient(Host, token)

	var offset int // смещение
	for {
		// получаем обновления (сообщения) от пользователя
		updates, err := client.Updates(offset, 10)
		if err != nil {
			log.Printf("error while getting updates: %v\n", err)
			time.Sleep(1 * time.Second) // Ждём секунду перед повторной попыткой в случае ошибки
			continue
		}

		// Обрабатываем каждое сообщение в цикле
		for _, update := range updates {
			// Выводим полученное сообщение в консоль (для отладки)
			fmt.Printf("New message from update %d: %v\n", update.Id, update.Message)
			// отправка ответа
			switch update.Message.Text {
			case "/start":
				StartParser(client, update.Message.Chat.Id, update.Message.User.FirstName)
				break
			default:
				answerForUser(client, int64(update.Message.Chat.Id), weather.GetWeather(update.Message.Text))
				// Обновляем offset, чтобы не получать старые сообщения повторно
			}
			offset = update.Id + 1

		}
		// Делаем паузу перед следующим запросом обновлений
		time.Sleep(2 * time.Second)
	}
}
