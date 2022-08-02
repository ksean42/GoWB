package main

import (
	_ "errors"
	"fmt"
	"io/ioutil"
	"log"
	_ "net/http"
	"os"

	"github.com/gin-gonic/gin"

	"encoding/json"

	"github.com/ksean42/GoWB/model"
	"github.com/ksean42/GoWB/server"
)

var Orders map[string]model.Order
var DB server.Database = server.Database{}

func main() {
	DB.DBInit()
	defer DB.DB.Close()
	Orders := server.RestoreCache(DB)
	server := server.NewServer()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
	sc, err := os.Open("m.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sc.Close()
	for k, v := range Orders {
		fmt.Printf("ID: %s", k)
		fmt.Println(v)
	}
	js, _ := ioutil.ReadAll(sc)

	var order model.Order
	json.Unmarshal(js, &order)

	DB.Create(order)
	DB.Get(order.OrderUid)

	router := gin.Default()
	router.GET("/examples/:id", getEx)
	router.GET("/examples/all", getExAll)
	router.POST("/examp", crEx)
	router.Run("localhost:8182")
}

func getEx(c *gin.Context) {
	id := c.Param("id")
	getedOrder, err := DB.Get(id)
	if err != nil {
		c.IndentedJSON(400, err)
	}
	c.IndentedJSON(200, getedOrder)
	// c.AbortWithStatus(http.StatusBadRequest)

}

func getExAll(c *gin.Context) {
	c.IndentedJSON(200, Orders)
	// c.AbortWithStatus(http.StatusBadRequest)

}
func crEx(c *gin.Context) {
	var newOrder model.Order
	if err := c.BindJSON(&newOrder); err != nil {
		fmt.Println("ASKDKDASSKADK")
		return
	}
	Orders[newOrder.OrderUid] = newOrder
	DB.Create(newOrder)
	c.IndentedJSON(201, Orders)
}
