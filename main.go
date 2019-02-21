package main

import (
	"fmt"

	"batman-server/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("I'm Batman!")

	g := gin.New()
	g.GET("/ws", lib.StartWS)
	g.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	g.Run(":8088")
}
