package main

import (
	"fmt"

	"brunomelo.crud/v1/routes"
)

func main() {
	fmt.Println("Running")
	routes.RouterHandlers()
}
