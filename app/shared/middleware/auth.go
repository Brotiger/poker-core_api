package middleware

import (
	"errors"
	"strings"

	"github.com/Brotiger/per-painted_poker-backend/app/shared/response"
	"github.com/gofiber/fiber/v2"
)

const headerPrefix = "Bearer"

func (am *AuthMiddleware) Token(c *fiber.Ctx) error {
	token, err := am.getTokenFromHeader(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error400{
			Message: "Неверный формат токена.",
		})
	}

	tokenClaims, err := am.tokenService.VerifyToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error401{
			Message: "Невалидный токен.",
		})
	}

	c.Locals("userId", tokenClaims.UserId)

	return c.Next()
}

func (am *AuthMiddleware) getTokenFromHeader(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization")
	l := len(headerPrefix)
	if len(header) < l+2 || header[:l] != headerPrefix {
		return "", errors.New("invalid token format")
	}

	return strings.TrimSpace(header[l:]), nil
}
