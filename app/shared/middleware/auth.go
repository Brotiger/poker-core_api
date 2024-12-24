package middleware

import (
	"github.com/Brotiger/per-painted_poker-backend/app/shared/response"
	"github.com/gofiber/fiber/v2"
)

func (a *Auth) Token(c *fiber.Ctx) error {
	token := c.GetRespHeader("Authorization")
	tokenClaims, err := a.ServiceToken.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error401{
			Message: "Невалидный токен.",
		})
	}

	c.Locals("userId", tokenClaims.UserId)

	return c.Next()
}
