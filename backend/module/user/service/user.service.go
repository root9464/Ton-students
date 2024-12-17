package user_service

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/root9464/Ton-students/backend/ent"
	user_repository "github.com/root9464/Ton-students/backend/module/user/repository"
	kafka_util "github.com/root9464/Ton-students/backend/shared/kafka"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	tma "github.com/telegram-mini-apps/init-data-golang"
)

func GetUserInDb(db *ent.Client, id int64) (*ent.User, error) {
	getUserInDb, err := user_repository.GetUserByID(db, id)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return getUserInDb, nil
}

func CreateUser(db *ent.Client, user *tma.InitData) (*ent.User, error) {
	createdUser, err := user_repository.CreateUser(db, user)

	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}

	return createdUser, nil
}

func ListenForUserRegistrations(db *ent.Client) {
	consumer, err := kafka_util.CreateKafkaConsumer("localhost:9092", "user_registrations", []string{"user_registrations"})
	if err != nil {
		log.Fatalf("Error creating Kafka Consumer: %v", err)
	}
	defer consumer.Close()

	topic := "user-register"
	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}

		var user *initdata.InitData
		err = json.Unmarshal(msg.Value, &user)
		if err != nil {
			log.Printf("Error unmarshaling user: %v", err)
			continue
		}

		_, err = CreateUser(db, user)
		if err != nil {
			log.Printf("Error creating user: %v", err)
			continue
		}

		log.Printf("Successfully created user: %v", user)
	}
}
