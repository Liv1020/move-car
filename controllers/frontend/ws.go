package frontend

import (
	"encoding/json"

	"net/http"

	"github.com/Liv1020/move-car-api/components"
	"github.com/Liv1020/move-car-api/middlewares"
	"github.com/Liv1020/move-car-api/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ws struct {
	Clients map[uint]*websocket.Conn
	Nexus   map[uint]map[uint]bool
}

// WS WS
var WS = ws{}

func init() {
	WS.Clients = make(map[uint]*websocket.Conn, 500)
	WS.Nexus = make(map[uint]map[uint]bool, 500)
}

// Handle Handle
func (t *ws) Handle(c *gin.Context) {
	auth := middlewares.JwtAuthFromClaims(c)
	db := components.App.DB()

	qr := c.Query("qr")
	if qr == "" {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	qrCode := new(models.Qrcode)
	if err := db.Preload("User").Where("id = ?", qr).Last(qrCode).Error; err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	upgrade := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		components.App.Logger().Errorf("Failed to set websocket upgrade: % v \n", err)
		return
	}

	// 加入客户端列表
	t.AddClient(auth.ID, conn)
	defer func() {
		t.RemoveClient(auth.ID)
	}()

	// 加入关系网
	t.AddNexus(qrCode.User.ID, auth.ID)
	defer func() {
		t.RemoveNexus(qrCode.User.ID, auth.ID)
	}()

	// 发送通知
	t.SendWait(qrCode.User.ID, qrCode.User.WaitMinute)

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		message := new(message)
		if err := json.Unmarshal(msg, message); err != nil {
			switch message.Type {
			case MessageTypeWait:
				auth.WaitMinute = message.Wait.Value
				if err := db.Save(auth).Error; err != nil {
					wsError(conn, mt, 1, err)
					continue
				}

				// 发送通知
				t.SendWait(qrCode.User.ID, message.Wait.Value)
			}
		}
	}
}

const (
	// MessageTypeWait MessageTypeWait
	MessageTypeWait = 1
)

type message struct {
	QrCode string `json:"qr_code"`
	Type   int    `json:"type"`
	Wait   struct {
		Value int `json:"value"`
	} `json:"wait"`
}

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func wsError(conn *websocket.Conn, messageType int, code int, err error) {
	response := new(response)
	response.Status = code
	response.Message = err.Error()
	by, err := json.Marshal(response)
	if err != nil {
		return
	}

	conn.WriteMessage(messageType, by)
}

func wsSuccess(conn *websocket.Conn, messageType int, data interface{}) {
	response := new(response)
	response.Status = 0
	response.Message = "Success"
	response.Data = data
	by, err := json.Marshal(response)
	if err != nil {
		return
	}

	conn.WriteMessage(messageType, by)
}

// SendWait SendWait
func (t *ws) SendWait(uid uint, wait int) {
	message := &message{
		Type: websocket.TextMessage,
		Wait: struct {
			Value int `json:"value"`
		}{
			Value: wait,
		},
	}

	// 发送通知
	if ns, ok := t.Nexus[uid]; ok {
		for uid := range ns {
			if c, ok := t.Clients[uid]; ok {
				wsSuccess(c, websocket.TextMessage, message)
			}
		}
	}
}

// AddClient AddClient
func (t *ws) AddClient(uid uint, conn *websocket.Conn) {
	t.RemoveClient(uid)
	t.Clients[uid] = conn
}

// RemoveClient RemoveClient
func (t *ws) RemoveClient(uid uint) {
	if client, ok := t.Clients[uid]; ok {
		client.Close()
		delete(t.Clients, uid)
	}
}

// AddNexus AddNexus
func (t *ws) AddNexus(owner uint, notice uint) {
	if _, ok := t.Nexus[owner]; !ok {
		t.Nexus[owner] = make(map[uint]bool, 1)
	}
	t.Nexus[owner][notice] = true
}

// RemoveNexus RemoveNexus
func (t *ws) RemoveNexus(owner uint, notice uint) {
	if len(t.Nexus[owner]) > 0 {
		delete(t.Nexus[owner], notice)
		if len(t.Nexus[owner]) == 0 {
			delete(t.Nexus, owner)
		}
	}
}
