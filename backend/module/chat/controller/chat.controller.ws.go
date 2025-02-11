package chat_controller

import (
	"context"
	"strconv"

	"github.com/gofiber/contrib/socketio"
	common_model "github.com/root9464/Ton-students/module/model/common"
)

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
		userID1, _ := strconv.ParseInt(ep.Kws.Params("id1"), 10, 64)
		userID2, _ := strconv.ParseInt(ep.Kws.Params("id2"), 10, 64)
		userIDs := []int64{userID1, userID2}
		ctx := context.Background()
		chatID, err := c.chatService.GetChatIDBetweenUsers(ctx, userIDs)
		if err != nil {
			c.socketErr(ep, err)
		}
	})

	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		userID, _ := strconv.ParseInt(ep.Kws.Params("id"), 10, 64)
		c.mu.Lock()
		defer c.mu.Unlock()
		delete(c.connections, userID)
		c.logger.Infof("User %v disconnected", userID)
	})

	socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {
		// userID, _ := strconv.ParseInt(ep.Kws.Params("id"), 10, 64)

		c.mu.RLock()
		defer c.mu.RUnlock()
		var chatMembers []common_model.ChatUser
		chatMembers = append(chatMembers, common_model.ChatUser{
			UserID: 123,
			ChatID: "1bc0ab1c-b334-4434-8ba4-3662053125f9",
		})

		chatMembers = append(chatMembers, common_model.ChatUser{
			UserID: 456,
			ChatID: "1bc0ab1c-b334-4434-8ba4-3662053125f9",
		})
		var membersConnections []string
		for _, member := range chatMembers {
			connection := c.connections[member.UserID]
			membersConnections = append(membersConnections, connection)
		}

		// ep.Kws.EmitToList(membersConnections, ep.Data)
		ep.Kws.Broadcast(ep.Data, false)
	})

	return func(kws *socketio.Websocket) {
		c.logger.Infof("KWS: %+v", kws)
	}
}
