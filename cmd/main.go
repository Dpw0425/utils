package main

import "utils/server"

func main() {
	app := server.InitServer()
	app.Run(":8080")
}
