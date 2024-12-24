package bot_service

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/root9464/Ton-students/shared/logger"
	"github.com/root9464/Ton-students/config"
)

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
func (s *BotService) Start(b *gotgbot.Bot, userID int64) error {
	// Логирование
	s.logger.Info(fmt.Sprintf("Start command received from user %d", userID))

	// Отправка сообщения пользователю
	_, err := b.SendMessage(userID, "Hello! How can I help you?", nil)
	if err != nil {
		s.logger.Error("Error sending start message: " + err.Error())
	}
	return err
}

// SupportStart - Логика обработки команды /support
func (s *BotService) SupportStart(b *gotgbot.Bot, userID int64, args []string) error {
	// Если нет аргументов, запрашиваем описание проблемы
	if len(args) == 0 {
		_, err := b.SendMessage(userID, "❓ <b>Введите ваш вопрос</b>\n\nПример:\n<code>/support Как зарегистрироваться?</code>", &gotgbot.SendMessageOpts{ParseMode: "HTML"})
		return err
	}

	// Если аргументы есть, обрабатываем их
	question := strings.Join(args, " ")
	s.logger.Info(fmt.Sprintf("Received support question from user %d: %s", userID, question))

	// Отправляем сообщение администратору
	_, err := b.SendMessage(s.config.AdminId, fmt.Sprintf("📩 <b>Новый запрос от пользователя</b>\n\n<b>Пользователь:</b> @%s\n<b>ID:</b> <code>%d</code>\n\n<b>Вопрос:</b>\n%s", userID, userID, question), &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		s.logger.Error("Error sending support message to admin: " + err.Error())
	}

	// Ответ пользователю
	_, err = b.SendMessage(userID, "✅ <b>Ваш запрос отправлен в поддержку</b>\n\nПожалуйста, ожидайте ответа от администратора.", &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}

// SupportReply - Логика для ответа на запросы поддержки
func (s *BotService) SupportReply(b *gotgbot.Bot, adminID int64, userID int64, messageText string) error {
	// Отправка сообщения пользователю
	_, err := b.SendMessage(userID, fmt.Sprintf("📬 <b>Ответ от администратора:</b>\n\n%s", messageText), &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	if err != nil {
		s.logger.Error("Error sending response to user: " + err.Error())
		return err
	}

	// Подтверждение администратора
	_, err = b.SendMessage(adminID, "✅ <b>Ваш ответ был успешно отправлен пользователю.</b>", &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}
