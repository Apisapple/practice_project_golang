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
		id := c.Param("id")
		result := db.GetMessage(id)
		c.IndentedJSON(http.StatusOK, result)
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
		var updateMessage model.Message

		if err := c.BindJSON(&updateMessage); err != nil {
			log.Fatal("BIND JSON ERROR")
			return
		}

		result := db.UpdateMessage(updateMessage)
		c.IndentedJSON(http.StatusAccepted, result)
	})

	router.DELETE("/message", func(c *gin.Context) {
		var deleteMessage model.Message

		if err := c.BindJSON(&deleteMessage); err != nil {
			log.Fatal("BIND JSON ERROR")
			return
		}

		result := db.DeleteMessage(deleteMessage.ID)
		c.IndentedJSON(http.StatusAccepted, result)
	})

	router.Run("localhost:8080")
}
