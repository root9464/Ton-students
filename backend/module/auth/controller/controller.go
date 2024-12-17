package auth_controller

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/config"
	"github.com/root9464/Ton-students/backend/ent"
	auth_service "github.com/root9464/Ton-students/backend/module/auth/service"
	kafka_util "github.com/root9464/Ton-students/backend/shared/kafka"
	"github.com/root9464/Ton-students/backend/shared/utils"
)

type Request struct {
	InitDataRaw string `json:"init_data_raw"`
}

func Authorize(log *utils.Logger, envs *config.Config, db *ent.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := new(Request)

		if err := ctx.BodyParser(request); err != nil {
			errResp := &fiber.Map{
				"message": err.Error(),
			}
			if err := ctx.Status(500).JSON(errResp); err != nil {
				log.Error("Failed to send error response:" + err.Error())
			}
			return err
		}

		expIn := 240 * time.Hour

		user, err := auth_service.ValidateInitData(request.InitDataRaw, envs.TelegramBotToken, expIn, log)
		if err != nil {
			errResp := &fiber.Map{
				"message": err.Error(),
			}
			if err := ctx.Status(500).JSON(errResp); err != nil {
				log.Error("Failed to send error response:" + err.Error())
			}
		}

		producer, err := kafka_util.CreateKafkaProducer("localhost:9092")
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"status":  "error",
				"message": "failed to create kafka producer",
			})
		}
		defer producer.Close()

		userData, err := json.Marshal(user)
		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"status":  "error",
				"message": "failed to marshal user data",
			})
		}

		topic := "user-register"
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          userData,
		}, nil)

		if err != nil {
			return ctx.Status(500).JSON(&fiber.Map{
				"status":  "error",
				"message": "failed to produce message",
			})
		}

		producer.Flush(15 * 1000)

		return ctx.Status(200).JSON(&fiber.Map{
			"status":  "success",
			"message": "create user",
			"user":    user,
		})
	}
}
