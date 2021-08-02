package main

import (
	"fmt"

	"github.com/ruwanego/scheddy"
	"github.com/ruwanego/scheddy/internal/auth"
	"github.com/ruwanego/scheddy/serverlib"
)

func main() {
	fmt.Println("Running server")
	fmt.Println("Config:", scheddy.Config())
	fmt.Println("Auth:", auth.GetAuth())
	fmt.Println(serverlib.Hello())
}
