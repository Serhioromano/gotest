package main

import (
	"fmt"

	"github.com/serhioromano/mygotest/internal/api"
	"github.com/serhioromano/mygotest/pkg/hello"

	"rsc.io/quote"
	quotev3 "rsc.io/quote/v3"
)

func main() {
	api.Hi("Sergey")
	fmt.Println(quote.Go())
	fmt.Println(quotev3.GoV3())
	fmt.Println(hello.Hello("Sergey"))
}
