package chat_controller

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	chat_service "github.com/root9464/Ton-students/module/chat/service"
	"github.com/root9464/Ton-students/shared/logger"
)

type ChatController struct {
	logger  *logger.Logger
	service *chat_service.ChatService
	rooms   map[string]map[*websocket.Conn]bool 
	mu      sync.Mutex
}

func NewChatController(logger *logger.Logger) *ChatController {
	return &ChatController{
		logger:  logger,
		service: chat_service.NewChatService(logger),
		rooms:   make(map[string]map[*websocket.Conn]bool),
	}
}

func (c *ChatController) HandleWebSocket(conn *websocket.Conn) {
	// Получаем ключ комнаты
	roomID := conn.Query("key")
	if roomID == "" || !c.service.RoomExists(roomID) {
		conn.WriteMessage(websocket.TextMessage, []byte("Error: Invalid room key"))
		conn.Close()
		return
	}

	c.mu.Lock()
	if _, exists := c.rooms[roomID]; !exists {
		c.rooms[roomID] = make(map[*websocket.Conn]bool)
	}
	c.rooms[roomID][conn] = true
	c.mu.Unlock()

	c.logger.Infof("User connected to room %s", roomID)

	defer func() {
		c.mu.Lock()
		delete(c.rooms[roomID], conn)
		c.mu.Unlock()
		c.logger.Infof("User disconnected from room %s", roomID)
		conn.Close()
	}()

	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			break 
		}

		c.logger.Infof("Message in room %s: %s", roomID, string(msg))

		c.mu.Lock()
		for client := range c.rooms[roomID] {
			if err := client.WriteMessage(messageType, msg); err != nil {
				client.Close()
				delete(c.rooms[roomID], client)
			}
		}
		c.mu.Unlock()
	}
}

func (c *ChatController) CreateRoomHandler(ctx *fiber.Ctx) error {
	roomID := c.service.CreateRoom()
	if roomID == "" {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create room"})
	}

	c.logger.Infof("Created new room: %s", roomID)
	return ctx.JSON(fiber.Map{"room_id": roomID})
}
