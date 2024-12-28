package chat_controller

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	chat_service "github.com/root9464/Ton-students/module/chat/service"
)

type ChatController struct {
	Hub         *chat_service.Hub
	ChatService *chat_service.ChatService
	Upgrader    *websocket.Upgrader
}

func NewChatController(hub *chat_service.Hub, chatService *chat_service.ChatService) *ChatController {
	return &ChatController{
		Hub:         hub,
		ChatService: chatService,
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (c *ChatController) HandleWebSocket(ctx *fiber.Ctx) error {
	// Преобразование fasthttp в стандартные net/http типы
	handler := func(w http.ResponseWriter, r *http.Request) {
		conn, err := c.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("WebSocket upgrade error:", err)
			return
		}

		client := &chat_service.Client{
			Hub:  c.Hub,
			Conn: conn,
			Send: make(chan []byte, 256),
		}

		// Регистрация клиента
		c.Hub.Register <- client

		// Запуск потоков для обработки клиента
		go client.WritePump()
		go client.ReadPump()
	}

	// Адаптируем запрос через fasthttpadaptor
	fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(handler))(ctx.Context())
	return nil
}