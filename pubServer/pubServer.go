package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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
		str := string(input)
		if err != nil {
			log.Println(err)
		}
		if str == "exit" {
			s.Close()
			os.Exit(0)
		}
		if str == "all" {
			publishAll(s)
			continue
		}
		publish(str, s)
	}

}

func publishAll(s stan.Conn) {
	files, err := ioutil.ReadDir("jsonTest/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
		if !file.IsDir() {
			publish(file.Name(), s)
			time.Sleep(time.Second * 2)
		}
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
