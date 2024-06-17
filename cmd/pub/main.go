package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func readJSON(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

const (
	clusterID = "test-cluster"
	clientID  = "publisher-client"
	url       = "nats://localhost:4222"
	subject   = "order_create"
)

func run() error {

	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		log.Fatalf("Failed to connect to NATS Streaming: %v", err)
	}
	defer sc.Close()

	jsonData := readJSON("orders.json")

	var orders []map[string]interface{}

	err = json.Unmarshal(jsonData, &orders)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	for _, order := range orders {
		bytes, err := json.Marshal(order)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
			continue
		}

		err = sc.Publish(subject, bytes)
		if err != nil {
			log.Fatalf("Error publishing message: %v", err)
		}
	}

	log.Println("All Messages published")

	return nil

}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

/*
Примеры тестовых данных в файле orders.json:

1: Валидный заказ
2: Тот же самый заказ
3: Полностью невалдиный заказ
4: Заказ с отрицательным  payment_dt
5: Заказ с неверной датой
6 и 7: Валидные заказы
8: отстуствет одно поле (oof_shard)
*/
