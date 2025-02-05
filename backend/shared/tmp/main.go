package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	USER_ID     = 99281932
	USERNAME    = "rogue"
	PRIVATE_KEY = "0x9de3bd5d174a7b4762cd4ca2d45680058054273daa5af4d568d6d616c597b438"
)

func generateAddressAndNickname(userID int, username, privateKey string) string {
	// Генерация данных для хеширования
	data := fmt.Sprintf("%d%s", userID, username)

	// Применяем HMAC-SHA256 с использованием privateKey
	mac := hmac.New(sha256.New, []byte(privateKey))
	mac.Write([]byte(data))
	hash := mac.Sum(nil)

	// Генерируем адрес в формате hex
	address := fmt.Sprintf("0x%s", hex.EncodeToString(hash))

	// Получаем никнейм из адреса
	addressWithoutPrefix := strings.TrimPrefix(address, "0x")
	nickname := make([]byte, 0, 8)

	// Преобразуем hex адрес в читаемый никнейм
	for i := 0; i < len(addressWithoutPrefix) && len(nickname) < 8; i++ {
		c := addressWithoutPrefix[i]
		var idx byte
		if c >= '0' && c <= '9' {
			idx = c - '0'
		} else {
			idx = 10 + c - 'a'
		}
		nickname = append(nickname, 'a'+idx)
	}

	// Возвращаем адрес и никнейм
	return string(nickname)
}

func main() {
	nickname := generateAddressAndNickname(USER_ID, USERNAME, PRIVATE_KEY)
	fmt.Println("Generated Nickname:", nickname)
}
