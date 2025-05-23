package main

import (
	"leaguesimulator/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080") // http://localhost:8080
}
