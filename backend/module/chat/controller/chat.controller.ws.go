package chat_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/contrib/socketio"
	chat_dto "github.com/root9464/Ton-students/module/chat/dto"
)

type MessageObject struct {
	Data  interface{} `json:"data"`
	Event string      `json:"event"`
}

func (c *ChatController) socketErr(ep *socketio.EventPayload, err error) {
	errByte := []byte(err.Error())
	ep.Kws.Emit(errByte)
}

func (c *ChatController) WS() func(*socketio.Websocket) {
	// var hub map[string][]string // hub connection by room ID
	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		userID, _ := strconv.ParseInt(ep.Kws.Params("id"), 10, 64)
		c.mu.Lock()
		defer c.mu.Unlock()
		c.connections[userID] = ep.Kws.UUID
		c.logger.Infof("User %v connected", userID)
	})

	socketio.On("join", func(ep *socketio.EventPayload) {
		message := new(MessageObject)
		if err := json.Unmarshal(ep.Data, message); err != nil {
			c.socketErr(ep, err)
			return
		}

		dataMap, ok := message.Data.(map[string]interface{})
		if !ok {
			c.socketErr(ep, nil)
			return
		}

		userIDFloat, ok := dataMap["user_id"].(float64)
		if !ok {
			c.socketErr(ep, nil)
			return
		}
		userIDStr := fmt.Sprintf("%.0f", userIDFloat) // Преобразуем float64 в строку
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			fmt.Println("Failed to parse user_id:", err)
			return
		}
		dto := chat_dto.CreateOrLoad{
			UserID:    userID,
			ServiceID: dataMap["service_id"].(string),
		}

		c.logger.Infof("join dto: %v", dto)
		ctx := context.Background()
		chatID, err := c.chatService.CreateOrLoadChat(ctx, &dto)
		if err != nil {
			c.socketErr(ep, err)
			return
		}

		c.logger.Infof("chatID: %v", chatID)
		c.userToChat[ep.Kws.UUID] = *chatID

		c.logger.Infof("userToChat: %v", c.userToChat)
		c.logger.Info("Join event success")
	})

	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		userID, _ := strconv.ParseInt(ep.Kws.Params("id"), 10, 64)
		c.mu.Lock()
		defer c.mu.Unlock()
		delete(c.connections, userID)
		c.logger.Infof("User %v disconnected", userID)
	})

	socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {
		message := new(MessageObject)
		if err := json.Unmarshal(ep.Data, message); err != nil {
			c.socketErr(ep, err)
		}

		if message.Event != "" {
			ep.Kws.Fire(message.Event, ep.Data)
		}
	})

	return func(kws *socketio.Websocket) {
		c.logger.Infof("KWS: %+v", kws)
	}
}
