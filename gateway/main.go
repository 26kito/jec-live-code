package main

import "gateway/routes"

func main() {
	app := routes.NewRouter()

	app.Listen(":8080")
}
