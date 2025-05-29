package main

import (
	"leaguesimulator/db"
	"leaguesimulator/routes"
	"log"
)

func main() {
	log.Println("Starting Advanced Football League Simulator...")
	db.InitDB()
	log.Println("Database connection is successful.")
	log.Println("Server will run on http://localhost:8080")
	log.Println("Available endpoints:")

	// Basic League Management
	log.Println("  GET /api/info - API information and features")
	log.Println("  POST /init-league - Initialize the league with 4 teams")
	log.Println("  POST /next-week - Play next week matches")
	log.Println("  POST /play-all - Play all remaining matches")
	log.Println("  POST /reset - Reset the league (with options)")

	// Data Retrieval
	log.Println("  GET /standings - Get current standings with enhanced info")
	log.Println("  GET /matches - Get all matches with statistics")
	log.Println("  GET /matches/week/:weekNumber - Get matches by specific week")
	log.Println("  GET /fixtures - Get future fixtures")

	// Predictions & Analytics
	log.Println("  GET /predict - Get comprehensive predictions with dynamic adjustments")
	log.Println("  GET /predict/:team1/:team2 - Get specific match prediction")
	log.Println("  GET /season-outlook - Get season outlook with elimination rules")
	log.Println("  GET /prediction-analytics - Get prediction model performance")

	// Team & League Analysis
	log.Println("  GET /team/:name/analysis - Get detailed team performance analysis")
	log.Println("  GET /league-stats - Get comprehensive league statistics")
	log.Println("  GET /head-to-head/:team1/:team2 - Get head-to-head comparison")

	// Match Management
	log.Println("  POST /edit-result - Edit match results with reason tracking")

	log.Println("\nFeatures:")
	log.Println("  • Multi-factor prediction algorithm")
	log.Println("  • Monte Carlo season simulation")
	log.Println("  • Dynamic elimination rules (Week 4: -6pts, Week 5: -3pts, Week 6: Winner takes all)")
	log.Println("  • Tactical analysis and weather impact modeling")
	log.Println("  • Team fatigue tracking and live match editing")
	log.Println("  • Championship probability calculation")
	log.Println("  • Advanced performance analytics")

	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
