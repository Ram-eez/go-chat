package main

import (
	"GO-CHAT/controllers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controllers.HomePage)
	r.GET("/ws", controllers.HandleConnections)

	go HandleMessages()
	fmt.Println("Server started on : 8080")
	err := r.Run(":8080")
	if err != nil {
		panic("Error starting server")
	}
}
