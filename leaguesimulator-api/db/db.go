package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	// Environment variables ile database bilgilerini al
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASSWORD", "emine")
	dbName := getEnv("DB_NAME", "leaguesimulator")

	// Production için güvenli DSN oluştur
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=30s&readTimeout=30s&writeTimeout=30s",
		dbUser, dbPass, dbHost, dbPort, dbName)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	// Production için optimize edilmiş connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(5 * time.Minute)
	DB.SetConnMaxIdleTime(time.Minute)

	// Connection test et
	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	log.Println("Database connection established successfully")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
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
