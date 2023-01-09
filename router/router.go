package router

import (
	"log"
	"strconv"

	db "example.com/practice/db"
	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/message/:id", func(c *gin.Context) {
		log.Fatal("Calling GET")
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)

		if err != nil {
			panic(err)
		}

		db.GetMessage(id)
	})

	router.POST("/message", func(c *gin.Context) {

	})

	router.Run(":8080")
}
