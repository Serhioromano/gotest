package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/serhioromano/mygotest/internal/utils"
)

type Inputs struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var EP = utils.EP{
	Runner:     getUsers,
	BodyParser: new(Inputs),
	Fields: map[string]utils.EPField{
		"ID": {
			Required: true,
			Type:     "number",
		},
		"Name": {
			Required: false,
			Type:     "string",
		},
	},
}

func init() {
	utils.EPRegister("usersList", EP)
}

func getUsers(c *fiber.Ctx) (map[string]interface{}, error) {
	fmt.Println("Users Lsist called")
	return map[string]interface{}{
		"name": "Sergey Romanov",
		"age":  100,
		"tel": map[string]interface{}{
			"mobile": "123456",
			"home":   "456789",
		},
	}, nil
}
