package middleware

import (
	"github.com/Brotiger/per-painted_poker-backend/internal/shared/response"
	"github.com/Brotiger/per-painted_poker-backend/internal/shared/service"
	"github.com/gofiber/fiber/v2"
)

var serviceToken *service.Token

func init() {
	serviceToken = service.NewToken()
}

func Token(c *fiber.Ctx) error {
	token := c.GetRespHeader("Authorization")
	tokenClaims, err := serviceToken.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error401{
			Message: "Невалидный токен.",
		})
	}

	c.Locals("userId", tokenClaims.UserId)

	return c.Next()
}
