package server

import (
	_ "fmt"
	"github.com/ksean42/GoWB/model"
	"log"
)

func RestoreCache(db Database) map[string]model.Order {
	Orders := make(map[string]model.Order)
	rows, err := db.GetAll()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	for rows.Next() {
		var uid string
		var order model.Order
		rows.Scan(&uid, &order)
		Orders[uid] = order
	}
	return Orders
}
