package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) (map[string]interface{}, error) {
	fmt.Println("Users Lsist called")
	return map[string]interface{}{
		"name":"Sergey Romanov",
		"age": 100,
		"tel": map[string]interface{}{
			"mobile": "123456",
			"home": "456789",
		},
	}, fiber.ErrBadRequest
}
