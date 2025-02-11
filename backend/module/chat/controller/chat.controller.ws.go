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
	From  string      `json:"from"`
	Event string      `json:"event"`
	To    string      `json:"to"`
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
		}

		dataMap, ok := message.Data.(map[string]interface{})
		if !ok {
			c.socketErr(ep, nil)
		}

		userIDFloat, ok := dataMap["user_id"].(float64)
		if !ok {
			c.socketErr(ep, nil)
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
		c.chatService.CreateOrLoadChat(ctx, &dto)

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

		// userID, _ := strconv.ParseInt(ep.Kws.Params("id"), 10, 64)
		// c.mu.RLock()
		// defer c.mu.RUnlock()
		// var chatMembers []common_model.ChatUser
		// chatMembers = append(chatMembers, common_model.ChatUser{
		// 	UserID: 123,
		// 	ChatID: "1bc0ab1c-b334-4434-8ba4-3662053125f9",
		// })
		//
		// chatMembers = append(chatMembers, common_model.ChatUser{
		// 	UserID: 456,
		// 	ChatID: "1bc0ab1c-b334-4434-8ba4-3662053125f9",
		// })
		// var membersConnections []string
		// for _, member := range chatMembers {
		// 	connection := c.connections[member.UserID]
		// 	membersConnections = append(membersConnections, connection)
		// }
		//
		// // ep.Kws.EmitToList(membersConnections, ep.Data)
		// ep.Kws.Broadcast(ep.Data, false)
	})

	return func(kws *socketio.Websocket) {
		c.logger.Infof("KWS: %+v", kws)
	}
}
