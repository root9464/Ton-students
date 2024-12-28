package config

import (
	"fmt"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type HTTPConfig interface {
	Address() string
	ChatAddres() string
}

type httpConfig struct {
	Host string
	Port string
	ChatPort string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("HTTP_HOST is not set")
	}
	port := os.Getenv(httpPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("HTTP_PORT is not set")
	}
	chatPort:= os.Getenv("CHAT_PORT")
	if chatPort == "" {
		return nil, fmt.Errorf("CHAT_PORT is not set")
	}
	return &httpConfig{
		Host: host,
		Port: port,
		ChatPort: chatPort,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.Host, cfg.Port)
}

func (cfg *httpConfig) ChatAddres() string {
	return net.JoinHostPort(cfg.Host, cfg.ChatPort)
}