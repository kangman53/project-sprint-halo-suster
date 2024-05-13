package helpers

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	exc "github.com/kangman53/project-sprint-halo-suster/exceptions"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func GetTokenHandler() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString("JWT_SECRET"))},
		ContextKey: JwtContextKey,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			auth := c.Locals(JwtContextKey).(*jwt.Token)
			claims := auth.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			c.Locals("userId", claims["user_id"].(string))
			c.Locals("userRole", claims["role"].(string))
			return c.Next()
		},
	})
}

func CheckTokenHeader(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return exc.UnauthorizedException("Unauthorized")
	} else {
		return ctx.Next()
	}
}
