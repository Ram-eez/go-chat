package controllers

import (
	"GO-CHAT/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HomePage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "welcome to the chat room"})

}

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}

	defer conn.Close()

	models.Clients[conn] = true

	for {
		var msg models.Message
		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Println(err)
			delete(models.Clients, conn)
			return
		}
		models.Broadcast <- msg
	}
}
