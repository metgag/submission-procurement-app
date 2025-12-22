package config

import (
	"os"
	"time"
)

var (
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	JWTExpiry = 12 * time.Hour
)
