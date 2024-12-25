package bot_service

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/root9464/Ton-students/config"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/sirupsen/logrus"
)

var replyState = struct {
	mu     sync.Mutex
	active map[int64]int64
}{
	active: make(map[int64]int64),
}

type BotService struct {
	config *config.Config
	logger *logger.Logger
}

func NewBotService(config *config.Config, logger *logger.Logger) *BotService {
	return &BotService{
		config: config,
		logger: logger,
	}
}

// Start - Логика обработки команды /start
func Start(b *gotgbot.Bot, ctx *ext.Context, channelid int64, log *logrus.Logger) error {
	userID := ctx.EffectiveUser.Id

	log.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("Start command received")

	member, err := b.GetChatMember(channelid, userID, nil)
	if err != nil {
		log.WithError(err).Error("Error checking subscription status")
		_, _ = ctx.EffectiveMessage.Reply(b, "Ошибка при проверке подписки. Попробуйте позже.", nil)
		return err
	}

	memberStatus := member.GetStatus()
	log.WithFields(logrus.Fields{
		"userID":       userID,
		"memberStatus": memberStatus,
	}).Info("Subscription status checked")

	if memberStatus != "member" && memberStatus != "administrator" && memberStatus != "creator" {
		_, err := ctx.EffectiveMessage.Chat.SendMessage(b,
			"Чтобы использовать этот сервис, пожалуйста, подпишитесь на канал [здесь](https://t.me/+M3Gb_96fCzNhMWFi)",
			nil,
		)
		if err != nil {
			log.WithError(err).Error("Error sending subscription requirement message")
		}
		return nil
	}

	_, err = b.SendMessage(ctx.EffectiveChat.Id, "Hello! How can I help you?", nil)
	if err != nil {
		log.WithError(err).Error("Error sending welcome message")
		return err
	}

	log.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("Welcome message sent successfully")

	return nil
}

// SupportStart - Логика обработки команды /support
func SupportStart(b *gotgbot.Bot, ctx *ext.Context, adminID int64, log *logrus.Logger) error {
	log.Info("SupportStart called by user ID:", ctx.EffectiveUser.Id)
	fmt.Println()

	userID := ctx.EffectiveUser.Id
	args := ctx.Args()

	if len(args) == 0 {
		_, err := ctx.EffectiveMessage.Reply(b,
			"❓ <b>Введите ваш вопрос</b>\n\nПример:\n<code>/support Как зарегистрироваться?</code>",
			&gotgbot.SendMessageOpts{ParseMode: "HTML"})
		return err
	}

	question := strings.Join(args, " ")
	log.WithFields(logrus.Fields{
		"userID":   userID,
		"question": question,
	}).Info("Received support question")

	_, err := b.SendMessage(adminID,
		fmt.Sprintf(
			"📩 <b>Новый запрос от пользователя</b>\n\n<b>Пользователь:</b> @%s\n<b>ID:</b> <code>%d</code>\n\n<b>Вопрос:</b>\n%s",
			ctx.EffectiveUser.Username, userID, question),
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = b.SendMessage(userID,
		"✅ <b>Ваш запрос отправлен в поддержку</b>\n\nПожалуйста, ожидайте ответа от администратора.",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Error(err)
		return err
	}

	replyMarkup := &gotgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
			{
				{
					Text:         "Ответить",
					CallbackData: fmt.Sprintf("reply_%d", userID),
				},
			},
		},
	}
	_, err = b.SendMessage(adminID,
		"👤 <b>Вы можете ответить на запрос этого пользователя:</b>",
		&gotgbot.SendMessageOpts{
			ParseMode:   "HTML",
			ReplyMarkup: replyMarkup,
		})
	return err
}

// SupportReply - Логика для ответа на запросы поддержки
func SupportReply(b *gotgbot.Bot, ctx *ext.Context, log *logrus.Logger) error {
	callbackData := ctx.CallbackQuery.Data
	log.Info("SupportReply called with callbackData:", callbackData)

	userIDStr := strings.TrimPrefix(callbackData, "reply_")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		log.Error("Error parsing userID from callbackData:", err)
		return err
	}

	replyState.mu.Lock()
	replyState.active[ctx.EffectiveUser.Id] = userID
	replyState.mu.Unlock()

	_, err = b.SendMessage(ctx.EffectiveUser.Id,
		"✍️ <b>Введите ваше сообщение для пользователя:</b>",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}

func SendAdminResponse(b *gotgbot.Bot, ctx *ext.Context, log *logrus.Logger) error {
	adminID := ctx.EffectiveUser.Id
	messageText := ctx.EffectiveMessage.Text
	fmt.Println(replyState.active)

	log.WithFields(logrus.Fields{
		"adminID": adminID,
		"message": messageText,
	}).Info("SendAdminResponse called")

	replyState.mu.Lock()
	userID, ok := replyState.active[adminID]
	if !ok {
		replyState.mu.Unlock()
		_, err := b.SendMessage(adminID,
			"⚠️ <b>Нет активного запроса для ответа.</b>\n\nСначала нажмите кнопку \"Ответить\".",
			&gotgbot.SendMessageOpts{ParseMode: "HTML"})
		return err
	}
	delete(replyState.active, adminID)
	replyState.mu.Unlock()

	_, err := b.SendMessage(userID,
		fmt.Sprintf(
			"📬 <b>Ответ от администратора:</b>\n\n%s",
			messageText),
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		log.Error(err)
		return err
	}

	_, err = b.SendMessage(adminID,
		"✅ <b>Ваш ответ был успешно отправлен пользователю.</b>",
		&gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}
