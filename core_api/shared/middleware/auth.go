package middleware

import (
	"errors"
	"strings"

	"github.com/Brotiger/poker-core_api/core_api/shared/response"
	"github.com/gofiber/fiber/v2"
)

const headerPrefix = "Bearer"

func (am *AuthMiddleware) Token(c *fiber.Ctx) error {
	token, err := am.getTokenFromHeader(c)
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

func (am *AuthMiddleware) getTokenFromHeader(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization")
	l := len(headerPrefix)
	if len(header) < l+2 || header[:l] != headerPrefix {
		return "", errors.New("invalid token format")
	}

	return strings.TrimSpace(header[l:]), nil
}
