package main

import (
	"encoding/json"
	"log"

	"github.com/ksean42/GoWB/model"
	"github.com/ksean42/GoWB/server"
	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "event-store1"
)

func main() {
	nats, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	defer nats.Close()

	s := server.NewServer()
	message := model.Order{}
	nats.Subscribe("order", func(msg *stan.Msg) {
		err = json.Unmarshal(msg.Data, &message)
		if err != nil {
			log.Println("Data format is wrong")
			return
		}
		if err := s.DB.Create(message); err != nil {
			log.Println(err)
		}
	})
	s.Start()
}
