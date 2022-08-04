package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ksean42/GoWB/model"
)

type Server struct {
	cache *Cache
	DB    Database
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	s.DB.DBInit()
	defer s.DB.CloseDb()
	s.cache = InitCache(s.DB)

	router := gin.Default()

	router.GET("/getOrder/:id", s.getOrder)
	router.GET("/getOrder/all", s.getAllOrders)
	router.POST("/order", s.saveOrder)
	router.Run("localhost:8282")
}

func (s *Server) getOrder(c *gin.Context) {
	id := c.Param("id")
	val, ok := s.cache.Orders[id]
	if ok {
		c.IndentedJSON(200, val)
		return
	}
	c.String(400, "Order doesn't exist")
}

func (s *Server) getAllOrders(c *gin.Context) {
	if len(s.cache.Orders) == 0 {
		c.String(500, "Database is empty")
		return
	}
	c.IndentedJSON(200, s.cache.Orders)
}

func (s *Server) saveOrder(c *gin.Context) {
	var newOrder model.Order

	if err := c.BindJSON(&newOrder); err != nil {
		c.String(400, "Something went wrong")
		return
	}
	s.cache.Orders[newOrder.OrderUid] = newOrder

	if err := s.DB.Create(newOrder); err != nil {
		c.IndentedJSON(400, err.Error())
		return
	}
	c.String(200, "Success")
}
