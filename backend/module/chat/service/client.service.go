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
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

// передает сообщения из вебсокета -> хаб
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)                                                                        //размер сообщения
	c.conn.SetReadDeadline(time.Now().Add(pongWait))                                                           // срок чтения
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil }) //обработка сообщений
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(message)
		c.hub.broadcast <- message
	}

}

// передает сообщения из хаба -> вебсокет
func (c *Client) writePump() {
	ticker := time.NewTicker(pingWait)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case massage, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(massage)
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
