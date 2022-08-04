package main

import (
	"github.com/ksean42/GoWB/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
