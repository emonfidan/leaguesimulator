# leaguesimulator
LeagueSimulator shows match results of a group of football teams, the league table and estimates the final league table.

# Football League Simulator - Setup Guide

## Project Structure
```
leaguesimulator/
├── main.go                 # Application entry point
├── go.mod                 # Go module dependencies
├── models/
│   └── models.go          # Data models for Team and Match
├── league/
│   └── leagueManager.go   # League management logic
├── prediction/
│   └── prediction.go      # AI prediction service
├── routes/
│   └── router.go          # API routes and handlers
├── prediction.py          # Python ML prediction script
└── README.md              # Documentation
```

## Prerequisites

1. **Go Installation** (version 1.21 or higher)
   ```bash
   # Download from https://golang.org/dl/
   go version  # Should show Go version
   ```

2. **Python Installation** (version 3.7 or higher)
   ```bash
   python3 --version  # or python --version
   ```

## Setup Instructions

1. **Clone/Create the project directory:**
   ```bash
   mkdir leaguesimulator
   cd leaguesimulator
   ```

2. **Initialize Go module:**
   ```bash
   go mod init leaguesimulator
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Make sure Python is accessible:**
   ```bash
   # Test Python command
   python3 --version
   # or
   python --version
   ```

5. **Run the application:**
   ```bash
   go run main.go
   ```

   You should see:
   ```
   Starting Football League Simulator...
   Server will run on http://localhost:8080
   Available endpoints:
     POST /init-league - Initialize the league
     ...
   ```

## API Testing Guide

### 1. Initialize League
```bash
curl -X POST http://localhost:8080/init-league
```
**Expected Response:**
```json
{
  "message": "League initialized with 4 teams",
  "teams": ["Lions", "Tigers", "Bears", "Wolves"],
  "season_structure": {
    "total_weeks": 3,
    "matches_per_week": 2,
    "total_matches": 6
  }
}
```

### 2. Play Next Week
```bash
curl -X POST http://localhost:8080/next-week
```
**Expected Response:**
```json
{
  "week": 1,
  "matches": [
    {
      "week": 1,
      "team1": "Lions",
      "team2": "Tigers", 
      "score1": 2,
      "score2": 1
    },
    {
      "week": 1,
      "team1": "Bears",
      "team2": "Wolves",
      "score1": 1,
      "score2": 0
    }
  ],
  "message": "Week completed successfully"
}
```

### 3. Get Current Standings
```bash
curl http://localhost:8080/standings
```
**Expected Response:**
```json
{
  "current_week": 1,
  "standings": [
    {
      "name": "Lions",
      "played": 1,
      "won": 1,
      "drawn": 0,
      "lost": 0,
      "goals_for": 2,
      "goals_against": 1,
      "goal_diff": 1,
      "points": 3
    }
    // ... other teams
  ],
  "total_teams": 4,
  "league_status": "ongoing"
}
```

### 4. Get AI Predictions
```bash
curl http://localhost:8080/predict
```
**Expected Response:**
```json
{
  "match_predictions": [
    {
      "team1": "Lions",
      "team2": "Tigers",
      "score1": 2,
      "score2": 1,
      "result": "Lions wins",
      "confidence": 75.5,
      "weather": "sunny",
      "expected_goals": {
        "home": 2.1,
        "away": 1.3
      },
      "win_probabilities": {
        "home_win": 65.2,
        "draw": 20.1,
        "away_win": 14.7
      }
    }
    // ... more predictions
  ],
  "season_simulation": {
    "championship_probabilities": {
      "Lions": 45.2,
      "Tigers": 28.7,
      "Bears": 15.1,
      "Wolves": 11.0
    },
    "simulation_runs": 1000
  }
}
```

### 5. Edit Match Result
```bash
curl -X POST http://localhost:8080/edit-result \
  -H "Content-Type: application/json" \
  -d '{
    "week": 1,
    "team1": "Lions",
    "team2": "Tigers",
    "score1": 3,
    "score2": 2,
    "reason": "Updated after video review"
  }'
```

### 6. Play All Remaining Matches
```bash
curl -X POST http://localhost:8080/play-all
```

### 7. Team Analysis
```bash
curl http://localhost:8080/team/Lions/analysis
```

### 8. Head-to-Head Comparison
```bash
curl http://localhost:8080/head-to-head/Lions/Tigers
```

### 9. League Statistics
```bash
curl http://localhost:8080/league-stats
```

### 10. Reset League
```bash
curl -X POST http://localhost:8080/reset
```

## Complete Testing Workflow

1. **Initialize the league:**
   ```bash
   curl -X POST http://localhost:8080/init-league
   ```

2. **Play all 3 weeks:**
   ```bash
   # Week 1
   curl -X POST http://localhost:8080/next-week
   
   # Week 2
   curl -X POST http://localhost:8080/next-week
   
   # Week 3
   curl -X POST http://localhost:8080/next-week
   ```

3. **Check final standings:**
   ```bash
   curl http://localhost:8080/standings
   ```

4. **Get predictions and analysis:**
   ```bash
   curl http://localhost:8080/predict
   curl http://localhost:8080/league-stats
   ```

5. **Test editing functionality:**
   ```bash
   curl -X POST http://localhost:8080/edit-result \
     -H "Content-Type: application/json" \
     -d '{"week": 1, "team1": "Lions", "team2": "Tigers", "score1": 4, "score2": 0}'
   ```

6. **Reset when done:**
   ```bash
   curl -X POST http://localhost:8080/reset
   ```

## Key Features

- **4 Teams:** Lions (90), Tigers (80), Bears (70), Wolves (60)
- **Round-Robin Format:** Each team plays each other once (3 weeks, 6 matches total)
- **Premier League Rules:** 3 points for win, 1 for draw, 0 for loss
- **AI Predictions:** Advanced ML predictions using Python
- **Match Editing:** Ability to edit match results and recalculate standings
- **Advanced Analytics:** Team analysis, head-to-head, league statistics
- **Season Simulation:** Monte Carlo simulation for championship probabilities

## Troubleshooting

### Python Issues
If prediction endpoints fail:
1. Check Python installation: `python3 --version` or `python --version`
2. Ensure `prediction.py` is in the project root
3. Check server logs for specific Python errors

### Port Issues
If port 8080 is in use:
```bash
# Kill process using port 8080
lsof -ti:8080 | xargs kill -9

# Or change port in main.go
r.Run(":8081")  // Use different port
```

### Module Issues
If Go modules fail:
```bash
go clean -modcache
go mod tidy
```

## API Documentation

All endpoints return JSON responses with proper HTTP status codes:
- **200 OK:** Successful operation
- **400 Bad Request:** Invalid request data
- **404 Not Found:** Resource not found
- **500 Internal Server Error:** Server-side error

The API supports CORS for frontend integration and includes comprehensive error handling and logging.