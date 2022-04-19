package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"

	witai "github.com/wit-ai/wit-go/v2"
)

var wolframClient *wolfram.Client

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {  //для каждого event в канале anlyticsChannel выводим информацию
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	godotenv.Load(".env")  //загружаем окружение

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))  //инициируем slacker бота
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))  //инициируем witai клиент
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}  //инициируем wolfram клиент
	go printCommandEvents(bot.CommandEvents())  //запускаем в отдельном потоке функцию printCommandEvents

	bot.Command("query for bot - <message>", &slacker.CommandDefinition{  //добавляем обработчик комманды по шаблону "query for bot - <message>"
		Description: "send any question to wolfram",  //добавляем описание
		Example:     "who is the president of india",  //добавляем пример
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {  //добавляем функцию обработки
			query := request.Param("message")  //получаем сообщение из запроса

			msg, _ := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			data, _ := json.MarshalIndent(msg, "", "    ")  //переводим message в json(каждый элемент начинается с "" и "    ")
			rough := string(data[:])
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value")
			answer := value.String()
			res, err := wolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("there is an error")
			}
			fmt.Println(value)
			response.Reply(res)  //посылаем ответ(res)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)  //запускаем бота

	if err != nil {
		log.Fatal(err)
	}
}