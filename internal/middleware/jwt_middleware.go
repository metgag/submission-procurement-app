package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/metgag/procurement-api-example/internal/config"
	"github.com/metgag/procurement-api-example/internal/utils"
)

func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		Claims: &utils.JWTClaims{},
		SigningKey: jwtware.SigningKey{
			Key: config.GetJWTSecret(),
		},
	})
}
