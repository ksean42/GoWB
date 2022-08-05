package server

import (
	_ "fmt"
	"log"
	"sync"

	"github.com/ksean42/GoWB/model"
)

type Cache struct {
	sync.Mutex
	Orders map[string]model.Order
}

func InitCache(db Database) *Cache {
	cache := &Cache{
		Orders: *RestoreCache(db),
	}
	return cache
}

func RestoreCache(db Database) *map[string]model.Order {
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
	return &Orders
}

func (c *Cache) Save(order *model.Order) {
	c.Lock()
	defer c.Unlock()
	c.Orders[order.OrderUid] = *order
}
