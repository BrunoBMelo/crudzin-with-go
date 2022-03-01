package main

import (
	"fmt"

	"brunomelo.crud/v1/internal/routes"
)

func main() {
	fmt.Println("Running")
	routes.RouterHandlers()
}
