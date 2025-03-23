package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Envs struct {
	OPEN_AI_KEY           string
	RPC_URL               string
	GAME_CONTRACT_ADDRESS string
	ADMIN_ADDRESS         string
	INIT_PROMPT           string
	PRIVATE_KEY           string
	PUBLIC_KEY            string
	WEB_PORT              string
	WEB_HOST              string
	USE_STATIC_KEYS       string
}

var loaded bool

var envs Envs

func LoadEnvs() Envs {
	if loaded {
		return envs
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	envs = Envs{
		OPEN_AI_KEY:           os.Getenv("OPEN_AI_KEY"),
		RPC_URL:               os.Getenv("RPC_URL"),
		GAME_CONTRACT_ADDRESS: os.Getenv("GAME_CONTRACT_ADDRESS"),
		ADMIN_ADDRESS:         os.Getenv("ADMIN_ADDRESS"),
		INIT_PROMPT:           os.Getenv("INIT_PROMPT"),
		PRIVATE_KEY:           os.Getenv("PRIVATE_KEY"),
		PUBLIC_KEY:            os.Getenv("PUBLIC_KEY"),
		WEB_PORT:              os.Getenv("WEB_PORT"),
		WEB_HOST:              os.Getenv("WEB_HOST"),
		USE_STATIC_KEYS:       os.Getenv("USE_STATIC_KEYS"),
	}
	loaded = true
	return envs
}
