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

func (c *ChatController) encodeMessage(ep *socketio.EventPayload) (map[string]interface{}, error) {
	message := new(MessageObject)
	if err := json.Unmarshal(ep.Data, message); err != nil {
		return nil, err
	}

	dataMap, ok := message.Data.(map[string]interface{})
	if !ok {
		c.socketErr(ep, nil)
		return nil, fmt.Errorf("failed to parse data")
	}

	return dataMap, nil
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
		dataMap, err := c.encodeMessage(ep)
		if err != nil {
			c.socketErr(ep, err)
			return
		}
		chatID, ok := dataMap["chat_id"].(string)
		if !ok {
			c.socketErr(ep, fmt.Errorf("failed to parse chat_id"))
			return
		}

		dto := chat_dto.Join{
			ChatID: chatID,
		}

		ctx := context.Background()
		_, err = c.chatService.GetChatByID(ctx, dto.ChatID)
		if err != nil {
			c.socketErr(ep, err)
			return
		}

		c.addUserToChat(ep.Kws.UUID, chatID)

		c.logger.Info("Join event success")
	})

	socketio.On("message-event", func(ep *socketio.EventPayload) {
		dataMap, err := c.encodeMessage(ep)
		if err != nil {
			c.socketErr(ep, err)
			return
		}

		chatID, ok := c.userToChat[ep.Kws.UUID]
		if !ok {
			c.socketErr(ep, fmt.Errorf("user not in chat"))
			return
		}

		participants, ok := c.chatToUsers[chatID]
		if !ok {
			c.socketErr(ep, fmt.Errorf("chat not found"))
			return
		}

		var recipients []string
		for _, uuid := range participants {
			if uuid != ep.Kws.UUID {
				recipients = append(recipients, uuid)
			}
		}

		message, ok := dataMap["message"].(string)
		if !ok {
			c.socketErr(ep, fmt.Errorf("failed to parse message"))
			return
		}

		ep.Kws.EmitToList(recipients, []byte(message))

		c.logger.Info("Message event success")
	})

	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		userID, _ := strconv.ParseInt(ep.Kws.Params("id"), 10, 64)
		c.mu.Lock()
		defer c.mu.Unlock()
		delete(c.connections, userID)
		c.removeUserFromChat(ep.Kws.UUID)
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

func (c *ChatController) addUserToChat(userUUID string, chatID string) {
	c.userToChat[userUUID] = chatID
	c.chatToUsers[chatID] = append(c.chatToUsers[chatID], userUUID)
}

func (c *ChatController) removeUserFromChat(userUUID string) {
	chatID, ok := c.userToChat[userUUID]
	if !ok {
		return // Пользователь не в чате
	}

	users := c.chatToUsers[chatID]
	for i, uuid := range users {
		if uuid == userUUID {
			c.chatToUsers[chatID] = append(users[:i], users[i+1:]...)
			break
		}
	}

	delete(c.userToChat, userUUID)
}
