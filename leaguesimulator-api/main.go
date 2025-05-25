package main

import (
	"leaguesimulator/db"
	"leaguesimulator/routes"
	"log"
)

func main() {
	log.Println("Starting Football League Simulator...")
	db.InitDB()
	log.Println("Database connection is successful.")
	log.Println("Server will run on http://localhost:8080")
	log.Println("Available endpoints:")
	log.Println("  POST /init-league - Initialize the league")
	log.Println("  POST /next-week - Play next week matches")
	log.Println("  GET /standings - Get current standings")
	log.Println("  GET /matches - Get all matches")
	log.Println("  GET /predict - Get AI predictions")
	log.Println("  GET /fixtures - Get future fixtures")
	log.Println("  POST /edit-result - Edit match results")
	log.Println("  POST /play-all - Play all remaining matches")
	log.Println("  POST /reset - Reset the league")

	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
