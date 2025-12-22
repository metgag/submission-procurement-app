package config

import (
	"os"
	"time"
)

var JWTExpiry = 12 * time.Hour

func GetJWTSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}
