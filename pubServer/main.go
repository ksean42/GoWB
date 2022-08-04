package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

const (
	clusterID = "test-cluster"
	clientID  = "event-store"
)

func main() {
	s, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter json filename to publish")
	for {
		input, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
		}
		if string(input) == "exit" {
			s.Close()
			os.Exit(0)
		}
		publish(string(input), s)
	}

}

func publish(filename string, s stan.Conn) {
	file, err := ioutil.ReadFile("jsonTest/" + filename)
	if err != nil {
		log.Println(err)
		return
	}
	if err := s.Publish("order", file); err != nil {
		log.Println(err)
	}
}
