package main

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/serhioromano/mygotest/configs"
	"github.com/serhioromano/mygotest/internal/utils"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	
	_ "github.com/serhioromano/mygotest/internal/users"
)

type Error struct {
	Success      bool   `json:"success"`
	ErrorCode    int16  `json:"code"`
	ErrorMessage string `json:"message"`
}

type Success struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

func SendError(i int, msg interface{}) Error {
	return Error{
		Success:      false,
		ErrorCode:    1,
		ErrorMessage: fmt.Sprint(msg),
	}
}

func main() {
	app := fiber.New(configs.FiberConfig)

	app.Post("/api/v1/:pkg/:function", func(c *fiber.Ctx) error {
		key := cases.Lower(language.AmericanEnglish).String(c.Params("pkg")) + cases.Title(language.AmericanEnglish).String(c.Params("function"))
		if handler := utils.EPRegestry[key]; handler.Runner != nil {
			err := c.BodyParser(handler.BodyParser)
			if err != nil {
				return c.JSON(SendError(500, err))
			}
			var v = reflect.ValueOf(handler.BodyParser)

			for key, val := range handler.Fields {
				value := reflect.Indirect(v).FieldByName(key)
				if !value.IsValid() {
					return c.JSON(SendError(500, "Cannot read field: "+key))
				}

				err := val.Validate(value)
				if err != nil {
					return c.JSON(SendError(500, err))
				}
			}

			data, err := handler.Runner(c)
			if err != nil {
				return c.JSON(SendError(500, err))
			}
			var out = Success{
				Success: true,
			}
			if data != nil {
				out.Data = data
			}
			fmt.Println(data)
			return c.JSON(out)
		}
		return fiber.ErrNotImplemented
	})

	app.Listen(":3000")
}
