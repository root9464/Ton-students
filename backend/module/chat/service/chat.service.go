package chat_service

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"sync"

	"github.com/root9464/Ton-students/shared/logger"
)

type Message struct {
	Data  string `json:"data"`
	From  string `json:"from"`
	Event string `json:"event"`
	To    string `json:"to"`
}

type ChatService struct {
	rooms map[string]bool
	mu    sync.RWMutex

	logger *logger.Logger
}

func NewChatService(logger *logger.Logger) *ChatService {
	return &ChatService{
		rooms:  make(map[string]bool),
		logger: logger,
	}
}

func (s *ChatService) CreateRoom() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		s.logger.Infof("Failed to generate private key: %v", err)
		return ""
	}

	roomID := hex.EncodeToString(privateKey.PublicKey.X.Bytes())
	s.rooms[roomID] = true
	return roomID
}

func (s *ChatService) RoomExists(roomID string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.rooms[roomID]
	return exists
}
