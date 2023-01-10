package router

import (
	"log"
	"net/http"

	db "example.com/practice/db"
	model "example.com/practice/model"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/message/:id", func(c *gin.Context) {
		log.Fatal("Calling GET")
		id := c.Param("id")
		db.GetMessage(id)
	})

	router.POST("/message", func(c *gin.Context) {
		var newMessage model.Message

		if err := c.BindJSON(&newMessage); err != nil {
			log.Fatal("BIND JSON ERROR")
			return
		}

		result := db.SaveMessage(newMessage)
		c.IndentedJSON(http.StatusCreated, result)
	})

	router.PUT("/message", func(c *gin.Context) {

	})

	router.DELETE("/message", func(c *gin.Context) {

	})

	router.Run("localhost:8080")
}
