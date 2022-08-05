package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksean42/GoWB/model"
)

type Server struct {
	Cache *Cache
	DB    Database
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	s.DB.DBInit()
	defer s.DB.CloseDb()
	s.Cache = InitCache(s.DB)

	router := gin.Default()
	router.LoadHTMLGlob("template/index.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})
	router.GET("/getOrder/:id", s.getOrder)
	router.GET("/getOrder/all", s.getAllOrders)
	router.POST("/order", s.saveOrder)
	router.Run("localhost:8181")
}

func (s *Server) getOrder(c *gin.Context) {
	id := c.Param("id")
	order, ok := s.Cache.Orders[id]
	if ok {
		c.IndentedJSON(200, order)
		return
	}
	c.IndentedJSON(400, "Order doesn't exist")
}

func (s *Server) getAllOrders(c *gin.Context) {
	if len(s.Cache.Orders) == 0 {
		c.IndentedJSON(500, "Database is empty")
		return
	}
	c.IndentedJSON(200, s.Cache.Orders)
}

func (s *Server) saveOrder(c *gin.Context) {
	var newOrder model.Order

	if err := c.BindJSON(&newOrder); err != nil {
		c.String(400, "Something went wrong")
		return
	}
	if err := Validate(&newOrder); err != nil {
		c.String(400, "Order is not valid")
		return
	}
	if err := s.DB.Create(&newOrder); err != nil {
		c.String(400, err.Error())
		return
	}
	s.Cache.Save(&newOrder)
	c.String(200, "Success")
}
