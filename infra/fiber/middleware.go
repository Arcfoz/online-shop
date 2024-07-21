package infrafiber

import (
	"fmt"
	"log"
	"online-shop/infra/response"
	"online-shop/internal/config"
	"online-shop/utility"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Trace() fiber.Handler {
	return func(c *fiber.Ctx) error {

		err := c.Next()

		return err
	}
}

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")
		// Auth : Bearer <token>
		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("token invalid")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		token := bearer[1]

		publicId, role, err := utility.ValidateToken(token, config.Cfg.App.Encrytion.JWTSecret)
		if err != nil {
			log.Println(err.Error())
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("ROLE", role)
		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}

func CheckRoles(autorizedRoles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role := fmt.Sprintf("%v", c.Locals("ROLE"))

		isExists := false
		for _, autorizedRoles := range autorizedRoles {
			if role == autorizedRoles {
				isExists = true
				break
			}
		}

		if !isExists {
			return NewResponse(
				WithError(response.ErrorForbidenAccess),
			).Send(c)
		}

		return c.Next()
	}
}
