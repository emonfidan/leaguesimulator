package main

import (
	"leaguesimulator/db"
	"leaguesimulator/routes"
	"log"
	"os"
)

func main() {
	log.Println("Starting Football League Simulator...")
	db.InitDB()
	log.Println("Database connection is successful.")
	log.Println("Server will run on http://localhost:8080")
	log.Println("Available endpoints:")
	log.Println("  GET /api/info - API information")
	log.Println("  POST /init-league - Initialize the league")
	log.Println("  POST /next-week - Play next week matches")
	log.Println("  GET /standings - Get current standings")
	log.Println("  GET /matches - Get all matches")
	log.Println("  GET /predict - Get predictions")
	log.Println("  GET /predict/:team1/:team2 - Get specific match prediction")
	log.Println("  GET /season-outlook - Get season outlook")
	log.Println("  GET /prediction-analytics - Get prediction analytics")
	log.Println("  GET /fixtures - Get future fixtures")
	log.Println("  POST /edit-result - Edit match results (legacy)")
	log.Println("  POST /play-all - Play all remaining matches")
	log.Println("  POST /reset - Reset the league")
	log.Println("  GET /team/:name/analysis - Get team analysis")
	log.Println("  GET /league-stats - Get league statistics")
	log.Println("  GET /head-to-head/:team1/:team2 - Get head-to-head comparison")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Lokal geliştirme için
	}

	r := routes.SetupRouter()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
