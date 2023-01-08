package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartRouter() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		log.Fatal("Start GO LANG Server~!!!")
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main Title",
		})
	})

	router.GET("/message", func(c *gin.Context) {

	})

	router.Run(":8080")
}
