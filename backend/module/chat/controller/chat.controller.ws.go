package chat_controller

import "github.com/gofiber/contrib/socketio"

func (c *ChatController) WS(kws *socketio.Websocket) {
	c.logger.Infof("KWS: %+v", kws)
}
