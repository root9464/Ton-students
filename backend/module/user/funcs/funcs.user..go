package user_funcs

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateUserNickname(userID int64, username, privateKey string) string {
	data := fmt.Sprintf("%d%s", userID, username)

	mac := hmac.New(sha256.New, []byte(privateKey))
	mac.Write([]byte(data))
	hash := mac.Sum(nil)

	address := fmt.Sprintf("0x%s", hex.EncodeToString(hash))
	addressWithoutPrefix := strings.TrimPrefix(address, "0x")
	nickname := make([]byte, 0, 8)

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

	return string(nickname)
}
