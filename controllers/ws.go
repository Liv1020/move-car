package controllers

import (
	"github.com/Liv1020/move-car/components"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ws struct {
}

// WS WS
var WS = ws{}

// Handle Handle
func (t *ws) Handle(c *gin.Context) {
	upgrade := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		components.App.Logger().Errorf("Failed to set websocket upgrade: % v \n", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
