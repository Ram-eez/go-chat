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

	// c.JSON(http.StatusOK, gin.H{"message": "welcome to the chat room"})
	c.String(200, "welcome to the chat room")

}

func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}

	defer conn.Close()

	models.Clients[conn] = true

	for {
		var msg models.Message
		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			delete(models.Clients, conn)
			return
		}
		models.Broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-models.Broadcast

		for client := range models.Clients {
			if err := client.WriteJSON(msg); err != nil {
				fmt.Println(err)
				client.Close()
				delete(models.Clients, client)
			}
		}
	}
}
