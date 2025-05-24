package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"leaguesimulator/league"
	"leaguesimulator/models"
	"leaguesimulator/prediction"
)

var manager league.LeagueManager
var predictionService *prediction.AdvancedPredictionService

func SetupRouter() *gin.Engine {
	router := gin.Default()
	predictionService = prediction.NewAdvancedPredictionService()

	// Enable CORS for frontend integration
	router.Use(corsMiddleware())

	// API Info endpoint
	router.GET("/api/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":        "Advanced Football League Simulator",
			"version":     "2.0.0",
			"description": "Enhanced league simulator with advanced ML predictions",
			"features": []string{
				"Multi-factor prediction algorithm",
				"Monte Carlo season simulation",
				"Tactical analysis",
				"Weather impact modeling",
				"Team fatigue tracking",
				"Live match editing",
				"Championship probability calculation",
			},
			"author": "Insider Internship Candidate",
		})
	})

	// Initialize league
	router.POST("/init-league", func(c *gin.Context) {
		manager.InitLeague()
		c.JSON(http.StatusOK, gin.H{
			"message": "League initialized with 4 teams",
			"teams":   []string{"Lions", "Tigers", "Bears", "Wolves"},
			"season_structure": gin.H{
				"total_weeks":      3,
				"matches_per_week": 2,
				"total_matches":    6,
			},
		})
	})

	// Play next week matches
	router.POST("/next-week", func(c *gin.Context) {
		matches := manager.PlayNextWeek()
		if matches == nil {
			c.JSON(http.StatusOK, gin.H{
				"message":         "League finished",
				"final_standings": manager.GetStandings(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"week":    manager.Week,
			"matches": matches,
			"message": "Week completed successfully",
		})
	})

	// Get current standings with enhanced info
	router.GET("/standings", func(c *gin.Context) {
		standings := manager.GetStandings()
		c.JSON(http.StatusOK, gin.H{
			"current_week": manager.Week,
			"standings":    standings,
			"total_teams":  len(standings),
			"league_status": func() string {
				if manager.Week >= 3 {
					return "completed"
				}
				return "ongoing"
			}(),
		})
	})

	// Get all matches with statistics
	router.GET("/matches", func(c *gin.Context) {
		matches := manager.GetMatches()

		// Calculate match statistics
		totalGoals := 0
		highestScore := 0
		for _, match := range matches {
			matchGoals := match.HomeGoals + match.AwayGoals
			totalGoals += matchGoals
			if matchGoals > highestScore {
				highestScore = matchGoals
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"matches": matches,
			"statistics": gin.H{
				"total_matches": len(matches),
				"total_goals":   totalGoals,
				"average_goals_per_match": func() float64 {
					if len(matches) == 0 {
						return 0.0
					}
					return float64(totalGoals) / float64(len(matches))
				}(),
				"highest_scoring_match": highestScore,
			},
		})
	})

	// Enhanced prediction endpoint
	router.GET("/predict", func(c *gin.Context) {
		predictions, err := predictionService.RunAdvancedPrediction()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to run advanced prediction: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, predictions)
	})

	// Specific match prediction
	router.GET("/predict/:team1/:team2", func(c *gin.Context) {
		team1 := c.Param("team1")
		team2 := c.Param("team2")

		prediction, err := predictionService.GetMatchPrediction(team1, team2)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Match prediction not found: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"match_prediction": prediction,
			"note":             "This is an AI-powered prediction based on multiple factors",
		})
	})

	// Season outlook endpoint
	router.GET("/season-outlook", func(c *gin.Context) {
		outlook, err := predictionService.GetSeasonOutlook()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get season outlook: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"season_outlook": outlook,
			"note":           "Based on 1000+ Monte Carlo simulations",
		})
	})

	// Prediction analytics
	router.GET("/prediction-analytics", func(c *gin.Context) {
		analytics := predictionService.CalculateAccuracy()
		c.JSON(http.StatusOK, gin.H{
			"analytics":   analytics,
			"description": "Prediction model performance metrics",
		})
	})

	// Play all remaining matches with detailed tracking
	router.POST("/play-all", func(c *gin.Context) {
		var allWeeks []gin.H
		weekCount := manager.Week

		for {
			matches := manager.PlayNextWeek()
			if matches == nil {
				break
			}
			weekCount++

			weekData := gin.H{
				"week":                 weekCount,
				"matches":              matches,
				"standings_after_week": manager.GetStandings(),
			}
			allWeeks = append(allWeeks, weekData)
		}

		finalStandings := manager.GetStandings()
		champion := ""
		if len(finalStandings) > 0 {
			champion = finalStandings[0].Name
		}

		c.JSON(http.StatusOK, gin.H{
			"message":         "All matches completed",
			"weeks_played":    allWeeks,
			"final_standings": finalStandings,
			"champion":        champion,
			"total_weeks":     len(allWeeks),
		})
	})

	// Enhanced reset with options
	router.POST("/reset", func(c *gin.Context) {
		var resetOptions struct {
			ResetType string `json:"reset_type"` // "full", "matches_only", "standings_only"
		}

		if err := c.ShouldBindJSON(&resetOptions); err != nil {
			// Default to full reset if no options provided
			resetOptions.ResetType = "full"
		}

		switch resetOptions.ResetType {
		case "full":
			manager.InitLeague()
		case "matches_only":
			manager.Matches = []models.Match{}
			manager.Week = 0
		case "standings_only":
			manager.UpdateStandings()
		default:
			manager.InitLeague()
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "League reset completed",
			"reset_type": resetOptions.ResetType,
			"current_state": gin.H{
				"week":           manager.Week,
				"matches_played": len(manager.GetMatches()),
				"teams":          len(manager.Teams),
			},
		})
	})

	// Get future fixtures with predictions
	router.GET("/fixtures", func(c *gin.Context) {
		fixtures := manager.GetFutureFixtures()
		c.JSON(http.StatusOK, gin.H{
			"upcoming_fixtures": fixtures,
			"total_remaining":   len(fixtures),
			"note":              "Use /predict endpoint for match predictions",
		})
	})

	// Enhanced edit match result
	router.POST("/edit-result", func(c *gin.Context) {
		var editRequest struct {
			Week   int    `json:"week" binding:"required"`
			Team1  string `json:"team1" binding:"required"`
			Team2  string `json:"team2" binding:"required"`
			Score1 int    `json:"score1" binding:"required"`
			Score2 int    `json:"score2" binding:"required"`
			Reason string `json:"reason,omitempty"`
		}

		if err := c.ShouldBindJSON(&editRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request format: " + err.Error(),
			})
			return
		}

		success := manager.EditMatchResult(
			editRequest.Week,
			editRequest.Team1,
			editRequest.Team2,
			editRequest.Score1,
			editRequest.Score2,
		)

		if !success {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Match not found or invalid parameters",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Match result updated successfully",
			"updated_match": gin.H{
				"week":      editRequest.Week,
				"teams":     editRequest.Team1 + " vs " + editRequest.Team2,
				"new_score": strconv.Itoa(editRequest.Score1) + "-" + strconv.Itoa(editRequest.Score2),
				"reason":    editRequest.Reason,
			},
			"updated_standings": manager.GetStandings(),
		})
	})

	// Team performance analysis
	router.GET("/team/:name/analysis", func(c *gin.Context) {
		teamName := c.Param("name")

		// Find team in current standings
		standings := manager.GetStandings()
		var teamData *league.TeamStanding
		position := 0

		for i, team := range standings {
			if team.Name == teamName {
				teamData = &standings[i]
				position = i + 1
				break
			}
		}

		if teamData == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error":           "Team not found",
				"available_teams": []string{"Lions", "Tigers", "Bears", "Wolves"},
			})
			return
		}

		// Calculate team-specific statistics
		matches := manager.GetMatches()
		teamMatches := []models.Match{}
		goalsInMatches := []int{}

		for _, match := range matches {
			if match.HomeTeam == teamName || match.AwayTeam == teamName {
				teamMatches = append(teamMatches, match)
				if match.HomeTeam == teamName {
					goalsInMatches = append(goalsInMatches, match.HomeGoals)
				} else {
					goalsInMatches = append(goalsInMatches, match.AwayGoals)
				}
			}
		}

		// Calculate form (last 3 matches)
		form := "N/A"
		if len(teamMatches) >= 3 {
			recentMatches := teamMatches[len(teamMatches)-3:]
			wins := 0
			draws := 0

			for _, match := range recentMatches {
				if match.HomeTeam == teamName {
					if match.HomeGoals > match.AwayGoals {
						wins++
					} else if match.HomeGoals == match.AwayGoals {
						draws++
					}
				} else {
					if match.AwayGoals > match.HomeGoals {
						wins++
					} else if match.AwayGoals == match.HomeGoals {
						draws++
					}
				}
			}

			if wins >= 2 {
				form = "Excellent"
			} else if wins == 1 && draws >= 1 {
				form = "Good"
			} else if draws >= 2 {
				form = "Average"
			} else {
				form = "Poor"
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"team":             teamName,
			"current_position": position,
			"performance_summary": gin.H{
				"points":          teamData.Points,
				"played":          teamData.Played,
				"won":             teamData.Won,
				"drawn":           teamData.Drawn,
				"lost":            teamData.Lost,
				"goals_for":       teamData.GoalsFor,
				"goals_against":   teamData.GoalsAgainst,
				"goal_difference": teamData.GoalDiff,
			},
			"advanced_stats": gin.H{
				"win_rate": func() float64 {
					if teamData.Played == 0 {
						return 0.0
					}
					return float64(teamData.Won) / float64(teamData.Played) * 100
				}(),
				"points_per_game": func() float64 {
					if teamData.Played == 0 {
						return 0.0
					}
					return float64(teamData.Points) / float64(teamData.Played)
				}(),
				"goals_per_game": func() float64 {
					if teamData.Played == 0 {
						return 0.0
					}
					return float64(teamData.GoalsFor) / float64(teamData.Played)
				}(),
				"goals_conceded_per_game": func() float64 {
					if teamData.Played == 0 {
						return 0.0
					}
					return float64(teamData.GoalsAgainst) / float64(teamData.Played)
				}(),
				"current_form": form,
			},
			"matches_played": teamMatches,
		})
	})

	// League statistics
	router.GET("/league-stats", func(c *gin.Context) {
		standings := manager.GetStandings()
		matches := manager.GetMatches()

		if len(standings) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "No league data available. Initialize league first.",
			})
			return
		}

		// Calculate league-wide statistics
		totalGoals := 0
		totalMatches := len(matches)
		highestScoringTeam := standings[0].Name
		bestDefense := standings[0].Name

		for _, team := range standings {
			totalGoals += team.GoalsFor
			if team.GoalsFor > standings[0].GoalsFor {
				highestScoringTeam = team.Name
			}
			for _, defTeam := range standings {
				if defTeam.GoalsAgainst < team.GoalsAgainst {
					bestDefense = defTeam.Name
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"league_overview": gin.H{
				"total_teams":    len(standings),
				"matches_played": totalMatches,
				"total_goals":    totalGoals,
				"average_goals_per_match": func() float64 {
					if totalMatches == 0 {
						return 0.0
					}
					return float64(totalGoals) / float64(totalMatches)
				}(),
				"current_week":         manager.Week,
				"league_leader":        standings[0].Name,
				"highest_scoring_team": highestScoringTeam,
				"best_defense":         bestDefense,
			},
			"standings": standings,
			"competition_status": func() string {
				if manager.Week >= 3 {
					return "Season Complete"
				}
				return "Season In Progress"
			}(),
		})
	})

	// Head-to-head comparison
	router.GET("/head-to-head/:team1/:team2", func(c *gin.Context) {
		team1 := c.Param("team1")
		team2 := c.Param("team2")

		matches := manager.GetMatches()
		h2hMatches := []models.Match{}
		team1Wins := 0
		team2Wins := 0
		draws := 0

		for _, match := range matches {
			if (match.HomeTeam == team1 && match.AwayTeam == team2) ||
				(match.HomeTeam == team2 && match.AwayTeam == team1) {
				h2hMatches = append(h2hMatches, match)

				if match.HomeTeam == team1 {
					if match.HomeGoals > match.AwayGoals {
						team1Wins++
					} else if match.HomeGoals < match.AwayGoals {
						team2Wins++
					} else {
						draws++
					}
				} else {
					if match.AwayGoals > match.HomeGoals {
						team1Wins++
					} else if match.AwayGoals < match.HomeGoals {
						team2Wins++
					} else {
						draws++
					}
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"teams": gin.H{
				"team1": team1,
				"team2": team2,
			},
			"head_to_head_record": gin.H{
				"matches_played": len(h2hMatches),
				"team1_wins":     team1Wins,
				"team2_wins":     team2Wins,
				"draws":          draws,
			},
			"matches": h2hMatches,
			"summary": func() string {
				if len(h2hMatches) == 0 {
					return "No matches played between these teams yet"
				}
				if team1Wins > team2Wins {
					return team1 + " leads the head-to-head record"
				} else if team2Wins > team1Wins {
					return team2 + " leads the head-to-head record"
				}
				return "Even head-to-head record"
			}(),
		})
	})

	return router
}

// CORS middleware for frontend integration
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
