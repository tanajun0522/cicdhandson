package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, makerGreeting(name))
	})
	router.Run(":8080")
}

func makerGreeting(name string) string {
	return fmt.Sprintf("Hello,%s!", name)
}
