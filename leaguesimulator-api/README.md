# League Simulator

LeagueSimulator shows match results of a group of football teams, the league table and estimates the final league table with advanced ML predictions and dynamic elimination rules.

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
│   └── prediction.go      # Advanced AI prediction service
├── routes/
│   └── router.go          # API routes and handlers
├── db/
│   └── database.go        # Database connection
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
   Starting Advanced Football League Simulator...
   Database connection is successful.
   Server will run on http://localhost:8080
   Available endpoints:
     GET /api/info - API information and features
     POST /init-league - Initialize the league with 4 teams
     ...
   ```

### Database Issues
  Required MySQL 8.0+
  Set environment variables in db/db.go:
    export DB_HOST=localhost
    export DB_PORT=3306
    export DB_USER=your_mysql_username
    export DB_PASSWORD=your_mysql_password
    export DB_NAME=leaguesimulator
    export PORT=8080
  Execute schema.sql

## API Testing Guide

### 1. API Information
```bash
curl http://localhost:8080/api/info
```
**Expected Response:**
```json
{
  "name": "Advanced Football League Simulator",
  "version": "2.0.0",
  "description": "Enhanced league simulator with advanced ML predictions",
  "features": [
    "Multi-factor prediction algorithm",
    "Monte Carlo season simulation",
    "Tactical analysis",
    "Weather impact modeling",
    "Team fatigue tracking",
    "Live match editing",
    "Championship probability calculation"
  ],
  "author": "Insider Internship Candidate"
}
```

### 2. Initialize League
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

### 3. Play Next Week
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
      "home_team": "Lions",
      "away_team": "Tigers", 
      "home_goals": 2,
      "away_goals": 1
    },
    {
      "week": 1,
      "home_team": "Bears",
      "away_team": "Wolves",
      "home_goals": 1,
      "away_goals": 0
    }
  ],
  "message": "Week completed successfully"
}
```

### 4. Get Enhanced Standings
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

### 5. Get Matches with Statistics
```bash
curl http://localhost:8080/matches
```
**Expected Response:**
```json
{
  "matches": [
    // ... match data
  ],
  "statistics": {
    "total_matches": 6,
    "total_goals": 15,
    "average_goals_per_match": 2.5,
    "highest_scoring_match": 5
  }
}
```

### 6. Get Matches by Week
```bash
curl http://localhost:8080/matches/week/1
```

### 7. Get Advanced Predictions (with Dynamic Adjustments)
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
      "predicted_score1": 2,
      "predicted_score2": 1,
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
  ],
  "season_simulation": {
    "championship_probabilities": {
      "Lions": 45.2,
      "Tigers": 28.7,
      "Bears": 15.1,
      "Wolves": 11.0
    },
    "simulation_runs": 1000,
    "methodology": "Adjusted probabilities: 1 team(s) eliminated based on week 5 rules"
  }
}
```

### 8. Get Specific Match Prediction
```bash
curl http://localhost:8080/predict/Lions/Tigers
```

### 9. Get Season Outlook (with Elimination Rules)
```bash
curl http://localhost:8080/season-outlook
```
**Expected Response:**
```json
{
  "season_outlook": {
    "championship_probabilities": {
      "Lions": 100.0,
      "Tigers": 0.0,
      "Bears": 0.0,
      "Wolves": 0.0
    },
    "methodology": "Adjusted probabilities: 3 team(s) eliminated based on week 6 rules"
  },
  "note": "Adjusted for week 6 elimination rules - Based on 1000+ Monte Carlo simulations",
  "current_week": 6,
  "standings_considered": 4
}
```

### 10. Get Prediction Analytics
```bash
curl http://localhost:8080/prediction-analytics
```

### 11. Team Performance Analysis
```bash
curl http://localhost:8080/team/Lions/analysis
```
**Expected Response:**
```json
{
  "team": "Lions",
  "current_position": 1,
  "performance_summary": {
    "points": 9,
    "played": 3,
    "won": 3,
    "drawn": 0,
    "lost": 0,
    "goals_for": 8,
    "goals_against": 2,
    "goal_difference": 6
  },
  "advanced_stats": {
    "win_rate": 100.0,
    "points_per_game": 3.0,
    "goals_per_game": 2.67,
    "goals_conceded_per_game": 0.67,
    "current_form": "Excellent"
  },
  "matches_played": [
    // ... team's match history
  ]
}
```

### 12. League Statistics
```bash
curl http://localhost:8080/league-stats
```
**Expected Response:**
```json
{
  "league_overview": {
    "total_teams": 4,
    "matches_played": 6,
    "total_goals": 15,
    "average_goals_per_match": 2.5,
    "current_week": 3,
    "league_leader": "Lions",
    "highest_scoring_team": "Lions",
    "best_defense": "Lions"
  },
  "standings": [
    // ... current standings
  ],
  "competition_status": "Season Complete"
}
```

### 13. Head-to-Head Comparison
```bash
curl http://localhost:8080/head-to-head/Lions/Tigers
```
**Expected Response:**
```json
{
  "teams": {
    "team1": "Lions",
    "team2": "Tigers"
  },
  "head_to_head_record": {
    "matches_played": 1,
    "team1_wins": 1,
    "team2_wins": 0,
    "draws": 0
  },
  "matches": [
    // ... head-to-head matches
  ],
  "summary": "Lions leads the head-to-head record"
}
```

### 14. Get Future Fixtures
```bash
curl http://localhost:8080/fixtures
```

### 15. Edit Match Result (Enhanced)
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
**Expected Response:**
```json
{
  "message": "Match result updated successfully",
  "updated_match": {
    "week": 1,
    "teams": "Lions vs Tigers",
    "new_score": "3-2",
    "reason": "Updated after video review"
  },
  "updated_standings": [
    // ... recalculated standings
  ]
}
```

### 16. Play All Remaining Matches
```bash
curl -X POST http://localhost:8080/play-all
```
**Expected Response:**
```json
{
  "message": "All matches completed",
  "weeks_played": [
    {
      "week": 1,
      "matches": [...],
      "standings_after_week": [...]
    }
    // ... more weeks
  ],
  "final_standings": [...],
  "champion": "Lions",
  "total_weeks": 3
}
```

### 17. Enhanced Reset with Options
```bash
# Full reset (default)
curl -X POST http://localhost:8080/reset

# Reset with specific options
curl -X POST http://localhost:8080/reset \
  -H "Content-Type: application/json" \
  -d '{"reset_type": "matches_only"}'
```
**Options:**
- `"full"` - Complete reset (default)
- `"matches_only"` - Reset only matches and week counter
- `"standings_only"` - Recalculate standings only

## Complete Testing Workflow

1. **Get API info:**
   ```bash
   curl http://localhost:8080/api/info
   ```

2. **Initialize the league:**
   ```bash
   curl -X POST http://localhost:8080/init-league
   ```

3. **Play all 3 weeks individually:**
   ```bash
   # Week 1
   curl -X POST http://localhost:8080/next-week
   
   # Week 2
   curl -X POST http://localhost:8080/next-week
   
   # Week 3
   curl -X POST http://localhost:8080/next-week
   ```

4. **Or play all at once:**
   ```bash
   curl -X POST http://localhost:8080/play-all
   ```

5. **Check comprehensive results:**
   ```bash
   curl http://localhost:8080/standings
   curl http://localhost:8080/league-stats
   curl http://localhost:8080/matches
   ```

6. **Get advanced predictions and analytics:**
   ```bash
   curl http://localhost:8080/predict
   curl http://localhost:8080/season-outlook
   curl http://localhost:8080/prediction-analytics
   ```

7. **Analyze specific teams:**
   ```bash
   curl http://localhost:8080/team/Lions/analysis
   curl http://localhost:8080/head-to-head/Lions/Tigers
   ```

8. **Test editing functionality:**
   ```bash
   curl -X POST http://localhost:8080/edit-result \
     -H "Content-Type: application/json" \
     -d '{"week": 1, "team1": "Lions", "team2": "Tigers", "score1": 4, "score2": 0, "reason": "Test edit"}'
   ```

9. **Reset when done:**
   ```bash
   curl -X POST http://localhost:8080/reset
   ```

## Key Features

### League Structure
- **4 Teams:** Lions, Tigers, Bears, Wolves
- **Round-Robin Format:** Each team plays each other once (3 weeks, 6 matches total)
- **Premier League Rules:** 3 points for win, 1 for draw, 0 for loss

### Advanced Prediction System
- **Multi-factor Algorithm:** Considers team strength, form, weather, fatigue
- **Monte Carlo Simulation:** 1000+ simulations for championship probabilities
- **Dynamic Elimination Rules:**
  - **Week 4+:** Teams more than 6 points behind are eliminated
  - **Week 5+:** Teams more than 3 points behind are eliminated
  - **Week 6+:** Only the highest-scoring team gets 100% probability

### Analytics & Features
- **Comprehensive Team Analysis:** Win rates, form, goals per game
- **Head-to-Head Records:** Direct comparison between any two teams
- **League Statistics:** Overall performance metrics and leaders
- **Live Match Editing:** Update results with reason tracking
- **Weather Impact Modeling:** Environmental factors in predictions
- **Tactical Analysis:** Advanced performance breakdowns

### Enhanced Functionality
- **Flexible Reset Options:** Full, matches-only, or standings-only reset
- **Week-specific Match Retrieval:** Get matches by specific week number
- **Future Fixtures:** See upcoming matches
- **Prediction Accuracy Tracking:** Monitor model performance
- **CORS Support:** Ready for frontend integration

## Dynamic Elimination System

The simulator implements a progressive elimination system:

1. **Weeks 1-3:** All teams remain eligible for championship
2. **Week 4:** Teams more than 6 points behind the leader are eliminated from championship calculations
3. **Week 5:** Teams more than 3 points behind the leader are eliminated
4. **Week 6:** Winner-takes-all - only the team with the highest points gets 100% championship probability

This system makes the season more exciting and realistic, reflecting how real competitions work.

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

The API supports CORS for frontend integration and includes comprehensive error handling, logging, and validation. All prediction endpoints use advanced machine learning algorithms with dynamic adjustments based on current league standings and elimination rules.