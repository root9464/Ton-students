package chat_service

import (
	"bytes"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingWait       = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

// передает сообщения из вебсокета -> хаб
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)                                                                        //размер сообщения
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))                                                           // срок чтения
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil }) //обработка сообщений
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(message)
		c.Hub.Broadcast <- message
	}

}

// передает сообщения из хаба -> вебсокет
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingWait)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case massage, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(massage)
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(<-c.Send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
