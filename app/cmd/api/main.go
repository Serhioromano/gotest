package main

import (
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
	"github.com/serhioromano/mygotest/configs"
	"github.com/serhioromano/mygotest/internal/users"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var EPDB = make(map[string]interface{})

type Error = struct {
	Success      bool   `json:"success"`
	ErrorCode    int16  `json:"code"`
	ErrorMessage string `json:"message"`
}

type Success = struct {
	Success bool                   `json:"success"`
	Data    map[string]interface{} `json:"data"`
}

func init() {
	list := users.Init()
	for key, value := range list {
		EPDB["users"+cases.Title(language.English).String(key)] = value
	}
}

func main() {
	app := fiber.New(configs.FiberConfig)

	app.Post("/api/v1/:pkg/:function", func(c *fiber.Ctx) error {
		callKey := cases.Lower(language.AmericanEnglish).String(c.Params("pkg")) + cases.Title(language.AmericanEnglish).String(c.Params("function"))

		if val, ok := EPDB[callKey]; ok {
			f := reflect.ValueOf(val)
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(c)

			result := f.Call(in)
			data := result[0].Interface()
			err := result[1].Interface()

			if err != nil {
				var out = Error{
					Success:      false,
					ErrorCode:    501,
					ErrorMessage: fmt.Sprint(err),
				}
				return c.JSON(out)
			}
			var out = Success{
				Success: true,
			}
			if data != nil {
				out.Data = data.(map[string]interface{})
			}
			fmt.Println(reflect.ValueOf(data).Kind())
			return c.JSON(out)
		}
		
		return fiber.ErrNotImplemented
	})

	app.Listen(":3000")
}
