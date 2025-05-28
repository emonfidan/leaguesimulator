package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Connection string with timeout and charset
	dsn := "root:emine@tcp(localhost:3306)/leaguesimulator?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Hour)

	// Test the connection
	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Database connection established successfully")
}

func GetHistoricalMatches() ([]map[string]interface{}, error) {
	query := `SELECT season, week, home_team_name, away_team_name, home_goals, away_goals 
              FROM historical_matches 
              ORDER BY season, week`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []map[string]interface{}
	for rows.Next() {
		var season, week, homeGoals, awayGoals int
		var homeTeam, awayTeam string

		err := rows.Scan(&season, &week, &homeTeam, &awayTeam, &homeGoals, &awayGoals)
		if err != nil {
			return nil, err
		}

		matches = append(matches, map[string]interface{}{
			"season":     season,
			"week":       week,
			"home_team":  homeTeam,
			"away_team":  awayTeam,
			"home_goals": homeGoals,
			"away_goals": awayGoals,
		})
	}
	return matches, nil
}
