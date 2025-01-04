package middleware

import (
	"github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/Brotiger/poker-core_api/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	tokenService *service.TokenService
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: service.NewTokenService(),
	}
}

func (am *AuthMiddleware) Token(c *fiber.Ctx) error {
	token, err := am.tokenService.GetToken(c.Get("Authorization"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.BadRequest{
			Message: "Неверный формат токена.",
		})
	}

	tokenClaims, err := am.tokenService.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Unauthorized{
			Message: "Невалидный токен.",
		})
	}

	c.Locals("userId", tokenClaims.UserId)

	return c.Next()
}
