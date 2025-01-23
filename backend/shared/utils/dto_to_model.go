package utils

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	user_model "github.com/root9464/Ton-students/module/user/model"
	"github.com/samber/lo"
)

func ConvertMapStructure[T, D any](dto D) (*T, error) {
	entity := new(T)
	config := &mapstructure.DecoderConfig{
		Result:     &entity,
		DecodeHook: mapstructure.StringToTimeHookFunc(time.RFC3339),
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(dto)
	return entity, err
}

func ConvertDtoToEntity[T, D any](dto D, opts ...copier.Option) (*T, error) {
	entity := new(T)
	err := copier.CopyWithOption(entity, dto, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func HandlerError(err error) (*fiber.Map, int) {
	if e, ok := err.(*fiber.Error); ok {
		return &fiber.Map{
			"status":  "error",
			"message": e.Message,
		}, e.Code
	}
	return nil, 0
}

func HexToKeys(privateKeyHex, publicKeyHex string) (ed25519.PrivateKey, ed25519.PublicKey, error) {
	privKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding private key: %w", err)
	}

	pubKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding public key: %w", err)
	}

	if len(privKeyBytes) != ed25519.PrivateKeySize {
		return nil, nil, fmt.Errorf("incorrect size of private key")
	}

	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return nil, nil, fmt.Errorf("incorrect size of public key")
	}

	return ed25519.PrivateKey(privKeyBytes), ed25519.PublicKey(pubKeyBytes), nil
}

func LimitSlice[T any](slice []T, maxLen int) []T {
	if len(slice) > maxLen {
		return slice[:maxLen]
	}
	return slice
}

func GetVisibleName(newUser *user_model.User) string {
	nameMap := lo.Assign(
		map[user_model.SelectedName]string{
			user_model.Username: newUser.Username,
		},
		lo.OmitBy(map[user_model.SelectedName]string{
			user_model.Username:  newUser.Username,
			user_model.Firstname: lo.FromPtr(newUser.Firstname),
			user_model.Lastname:  lo.FromPtr(newUser.Lastname),
			user_model.Nickname:  lo.FromPtr(newUser.Nickname),
		}, func(_ user_model.SelectedName, value string) bool {
			return value == ""
		}),
	)

	entry, found := lo.Find(lo.Entries(nameMap), func(entry lo.Entry[user_model.SelectedName, string]) bool {
		return entry.Key == newUser.SelectedName
	})

	if found {
		return entry.Value
	}

	return "none"
}
