package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	jwt_dto "github.com/root9464/Ton-students/module/jwt/dto"
	jwt_funcs "github.com/root9464/Ton-students/module/jwt/functions"
	jwt_helpers "github.com/root9464/Ton-students/module/jwt/helpers"
	"github.com/root9464/Ton-students/shared/logger"
)

// var (
// 	// Приватный и публичный ключи Ed25519
// 	privateKey ed25519.PrivateKey
// 	publicKey  ed25519.PublicKey
// )

// type UserData struct {
// 	ID        uint64
// 	Username  string
// 	FirstName string
// 	LastName  *string
// }

// func init() {
// 	// Генерация пары ключей
// 	var err error
// 	publicKey, privateKey, err = ed25519.GenerateKey(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func createJWT(claims jwt.Claims, key ed25519.PrivateKey) (string, error) {
// 	// Создание нового JWT токена
// 	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

// 	// Подпись токена приватным ключом
// 	signedToken, err := token.SignedString(key)
// 	if err != nil {
// 		return "", err
// 	}

// 	return signedToken, nil
// }

// func verifyJWT(tokenString string, key ed25519.PublicKey) (*jwt.Token, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// Возвращаем публичный ключ для проверки подписи
// 		return key, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }

// func hashSalt(salt string) string {
// 	hash := sha256.New()
// 	hash.Write([]byte(salt))
// 	return hex.EncodeToString(hash.Sum(nil))
// }

// func main() {
// 	// Инициализация пользователя (соль)
// 	user := UserData{
// 		ID:        2200607311,
// 		Username:  "rootton_vf",
// 		FirstName: "RootTon",
// 		LastName:  nil,
// 	}

// 	rawUser := fmt.Sprintf("%s%s%d", user.Username, user.FirstName, user.ID)

// 	hashedSalt := hashSalt(rawUser)

// 	// Вывод ключей
// 	fmt.Println("Private Key:", hex.EncodeToString(privateKey))
// 	fmt.Println("Public Key:", hex.EncodeToString(publicKey))

// 	// Создание claims для JWT с добавлением информации о пользователе
// 	accessClaims := jwt.MapClaims{
// 		"sub":       "1234567890",                            // Subject (например, ID пользователя)
// 		"iat":       time.Now().Unix(),                       // Время выпуска токена
// 		"exp":       time.Now().Add(15 * time.Minute).Unix(), // Время истечения токена (15 минут)
// 		"user_hash": hashedSalt,
// 	}

// 	refreshClaims := jwt.MapClaims{
// 		"sub":       "1234567890",
// 		"iat":       time.Now().Unix(),
// 		"exp":       time.Now().Add(24 * time.Hour).Unix(), // Время истечения refresh токена (24 часа)
// 		"user_hash": hashedSalt,
// 	}

// 	// Создание и подписание JWT токенов
// 	accessToken, err := createJWT(accessClaims, privateKey)
// 	if err != nil {
// 		log.Fatal("Error creating access token:", err)
// 	}

// 	refreshToken, err := createJWT(refreshClaims, privateKey)
// 	if err != nil {
// 		log.Fatal("Error creating refresh token:", err)
// 	}

// 	// Вывод подписанных токенов
// 	fmt.Println("Access Token:", accessToken)
// 	fmt.Println("Refresh Token:", refreshToken)

// 	// Проверка и верификация токенов с использованием публичного ключа
// 	// Проверка access токена
// 	verifiedAccessToken, err := verifyJWT(accessToken, publicKey)
// 	if err != nil {
// 		log.Fatal("Error verifying access token:", err)
// 	}
// 	fmt.Println("Access token is valid:", verifiedAccessToken.Valid)

// 	// Проверка refresh токена
// 	verifiedRefreshToken, err := verifyJWT(refreshToken, publicKey)
// 	if err != nil {
// 		log.Fatal("Error verifying refresh token:", err)
// 	}
// 	fmt.Println("Refresh token is valid:", verifiedRefreshToken.Valid)

// }

const (
	privateKey = "38b87a8b0cf0ea5321324a35b8a7dfa656201cba91f2ac305a2202970b0a5e735d6db54cbc188e707e0be32c3f4c54b5f5984c7fdd7f81201d8cd0cbf321a0e5"
	publicKey  = "5d6db54cbc188e707e0be32c3f4c54b5f5984c7fdd7f81201d8cd0cbf321a0e5"
)

func hexToKeys(privateKeyHex, publicKeyHex string) (ed25519.PrivateKey, ed25519.PublicKey, error) {
	// Декодирование hex строк в байты
	privKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка декодирования private key: %w", err)
	}

	pubKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка декодирования public key: %w", err)
	}

	// Проверка на соответствие размерности ключей
	if len(privKeyBytes) != ed25519.PrivateKeySize {
		return nil, nil, fmt.Errorf("неправильный размер private key")
	}

	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return nil, nil, fmt.Errorf("неправильный размер public key")
	}

	// Возвращаем private и public ключи
	return ed25519.PrivateKey(privKeyBytes), ed25519.PublicKey(pubKeyBytes), nil
}

func main() {

	privK, pubK, err := hexToKeys(privateKey, publicKey)
	if err != nil {
		panic(err)
	}

	helpers := jwt_helpers.NewJwtHelper(logger.GetLogger(), validator.New())

	jwt := jwt_funcs.NewJwtFuncs(
		logger.GetLogger(),
		validator.New(),
		privK,
		pubK,
		helpers,
	)

	accessToken, refreshToken, _ := jwt.GenerateKeyPair(jwt_dto.UserData{
		ID:        2200607311,
		Username:  "rootton_vf",
		FirstName: "RootTon",
		LastName:  nil,
	})

	verifiedAccessToken, err := helpers.VerifyJwt(*accessToken, pubK)
	if err != nil {
		log.Fatal("Error verifying access token:", err)
	}
	fmt.Println("Access token is valid:", verifiedAccessToken.Valid)

	// Проверка refresh токена
	verifiedRefreshToken, err := helpers.VerifyJwt(*refreshToken, pubK)
	if err != nil {
		log.Fatal("Error verifying refresh token:", err)
	}
	fmt.Println("Refresh token is valid:", verifiedRefreshToken.Valid)

	fmt.Println("Access Token:", *accessToken)
	fmt.Println("Refresh Token:", *refreshToken)

}
