package utils

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type EPRunner func(c *fiber.Ctx) (map[string]interface{}, error)

type EPField struct {
	Required bool
	Type string // what is type string or URL or phone number
}
func (epf *EPField) Validate(val interface{}) error {
	str := fmt.Sprintf("%v", val);
	if epf.Required && (str == "" || str == "0") {
		return errors.New("field && is required")
	}
	
	fmt.Println(str)
	switch epf.Type {
	case "num":

	}
	return nil
}

type EP struct {
    Fields map[string]EPField
	BodyParser interface{}
    Runner  EPRunner
}

var EPRegestry = make(map[string]EP)

func EPRegister(name string, f EP) {
	EPRegestry[name] = f
}
