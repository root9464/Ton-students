package service

import (
	"fmt"
	"strings"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/sirupsen/logrus"
	"root/pkg/config"
)

type BotService struct {
	bot    *gotgbot.Bot
	config *config.Config
	log    *logrus.Logger
}

var replyState = struct {
	mu     sync.Mutex
	active map[int64]int64
}{
	active: make(map[int64]int64),
}

func NewBotService(bot *gotgbot.Bot, config *config.Config, log *logrus.Logger) *BotService {
	return &BotService{
		bot:    bot,
		config: config,
		log:    log,
	}
}

// Start command logic
func (s *BotService) Start(b *gotgbot.Bot, ctx *ext.Context) error {
	userID := ctx.EffectiveUser.Id

	// Check if the user is subscribed
	member, err := b.GetChatMember(s.config.ChannelId, userID, nil)
	if err != nil {
		s.log.WithError(err).Error("Error checking subscription")
		_, _ = ctx.EffectiveMessage.Reply(b, "Ошибка при проверке подписки. Попробуйте позже.", nil)
		return err
	}

	if member.GetStatus() != "member" && member.GetStatus() != "administrator" && member.GetStatus() != "creator" {
		_, _ = ctx.EffectiveMessage.Chat.SendMessage(b, "Please subscribe to the channel.", nil)
		return nil
	}

	_, err = b.SendMessage(ctx.EffectiveChat.Id, "Hello! How can I help you?", nil)
	return err
}

// Support command logic
func (s *BotService) SupportStart(b *gotgbot.Bot, ctx *ext.Context) error {
	args := ctx.Args()
	if len(args) == 0 {
		_, err := ctx.EffectiveMessage.Reply(b, "Please enter your question.", nil)
		return err
	}

	question := strings.Join(args, " ")
	userID := ctx.EffectiveUser.Id

	// Send the support request to the admin
	_, err := b.SendMessage(s.config.AdminId, fmt.Sprintf("New support request from user %s: %s", ctx.EffectiveUser.Username, question))
	if err != nil {
		s.log.Error("Error sending support request:", err)
		return err
	}

	// Reply to the user
	_, err = b.SendMessage(userID, "Your request has been sent to support.")
	return err
}

// Reply to user query
func (s *BotService) SupportReply(b *gotgbot.Bot, ctx *ext.Context) error {
	userIDStr := strings.TrimPrefix(ctx.CallbackQuery.Data, "reply_")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		s.log.Error("Error parsing userID:", err)
		return err
	}

	replyState.mu.Lock()
	replyState.active[ctx.EffectiveUser.Id] = userID
	replyState.mu.Unlock()

	_, err = b.SendMessage(ctx.EffectiveUser.Id, "Please enter your response to the user.")
	return err
}

// Send response to the user
func (s *BotService) SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context) error {
	adminID := ctx.EffectiveUser.Id
	messageText := ctx.EffectiveMessage.Text

	replyState.mu.Lock()
	userID, ok := replyState.active[adminID]
	if !ok {
		replyState.mu.Unlock()
		_, err := b.SendMessage(adminID, "No active request to reply to.")
		return err
	}
	delete(replyState.active, adminID)
	replyState.mu.Unlock()

	_, err := b.SendMessage(userID, fmt.Sprintf("Admin response: %s", messageText))
	if err != nil {
		s.log.Error("Error sending response to user:", err)
		return err
	}

	_, err = b.SendMessage(adminID, "Your response has been sent to the user.")
	return err
}
