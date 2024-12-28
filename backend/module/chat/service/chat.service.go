package chat_service

import "log"

type ChatService struct {
	Hub *Hub
}

func NewChatService(h *Hub) *ChatService {
	return &ChatService{Hub: h}
}

// HandleMessage обрабатывает сообщение, отправленное пользователем.
func (h *Hub) HandleMessage(client *Client, message []byte) {
	log.Printf("отправленное сообщение: %s", message)
	h.Broadcast <- message
}

// HandleConnection обраба
func (h *Hub) HandleConnection(client *Client) {
	log.Printf("клиент подключился: %s", client.Conn.RemoteAddr().String())
	h.Clients[client] = true
}

// Disconnect обрабатывает отключение клиента.
func (h *Hub) HandleDisconnect(client *Client) {
	log.Printf("клиент отключился: %s", client.Conn.RemoteAddr().String())
	h.Clients[client] = false
}
