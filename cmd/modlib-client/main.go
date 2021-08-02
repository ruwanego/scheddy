package main

import (
	"fmt"

	"github.com/ruwanego/scheddy"
	"github.com/ruwanego/scheddy/clientlib"
)

func main() {
	fmt.Println("Running client")
	fmt.Println("Config:", scheddy.Config())
	fmt.Println(clientlib.Hello())
}
