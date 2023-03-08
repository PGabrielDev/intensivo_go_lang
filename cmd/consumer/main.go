package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github/pgabrieldeveloper/intensivo_go/internal/infra/database"
	usecases "github/pgabrieldeveloper/intensivo_go/internal/use_cases"
	"github/pgabrieldeveloper/intensivo_go/pkg/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "modernc.org/sqlite"
)

func main() {

	db, err := sql.Open("sqlite", "./orders.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)

	usecase := usecases.CalculateFinalPrice{OrderRepository: repository}
	msgKafkaChan := make(chan *ckafka.Message)
	topics := []string{"orders"}
	servers := "host.docker.internal:9094"
	go kafka.Consume(topics, servers, msgKafkaChan)
	kafkaWorker(msgKafkaChan, usecase)
}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecases.CalculateFinalPrice) {
	for msg := range msgChan {
		//processa menssagem
		var orderInput usecases.OrderInputDTO
		err := json.Unmarshal(msg.Value, &orderInput)
		if err != nil {
			panic(err)
		}

		orderOutPut, err := uc.Execute(orderInput)

		if err != nil {
			panic(err)
		}
		fmt.Println(msg.Key, "Essa a key da mensagem")
		fmt.Printf("Kafka has processed order %s \n", orderOutPut.ID)

	}
}
