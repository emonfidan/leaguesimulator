package prediction

import (
	"encoding/json"
	"fmt"
	"leaguesimulator/db"
	"log"
	"os/exec"
	"runtime"
	"time"
)

type MatchPrediction struct {
	Team1            string           `json:"team1"`
	Team2            string           `json:"team2"`
	Score1           int              `json:"score1"`
	Score2           int              `json:"score2"`
	Result           string           `json:"result"`
	Confidence       float64          `json:"confidence"`
	Weather          string           `json:"weather"`
	MatchImportance  string           `json:"match_importance"`
	HomeStrength     float64          `json:"home_strength"`
	AwayStrength     float64          `json:"away_strength"`
	ExpectedGoals    ExpectedGoals    `json:"expected_goals"`
	WinProbabilities WinProbabilities `json:"win_probabilities"`
	TacticalAnalysis TacticalAnalysis `json:"tactical_analysis"`
}

type ExpectedGoals struct {
	Home float64 `json:"home"`
	Away float64 `json:"away"`
}

type WinProbabilities struct {
	HomeWin float64 `json:"home_win"`
	Draw    float64 `json:"draw"`
	AwayWin float64 `json:"away_win"`
}

type TacticalAnalysis struct {
	KeyBattles          []string          `json:"key_battles"`
	TacticalAdvantages  map[string]string `json:"tactical_advantages"`
	RecommendedStrategy map[string]string `json:"recommended_strategy"`
}

type SeasonSimulation struct {
	ChampionshipProbabilities map[string]float64 `json:"championship_probabilities"`
	SimulationRuns            int                `json:"simulation_runs"`
	Methodology               string             `json:"methodology"`
}

type PredictionMetadata struct {
	Algorithm         string   `json:"algorithm"`
	FactorsConsidered []string `json:"factors_considered"`
	ConfidenceLevel   string   `json:"confidence_level"`
	LastUpdated       string   `json:"last_updated"`
}

type ComprehensivePrediction struct {
	MatchPredictions   []MatchPrediction  `json:"match_predictions"`
	SeasonSimulation   SeasonSimulation   `json:"season_simulation"`
	PredictionMetadata PredictionMetadata `json:"prediction_metadata"`
}

// PredictionService interface for better architecture
type PredictionService interface {
	RunAdvancedPrediction() (*ComprehensivePrediction, error)
	GetMatchPrediction(team1, team2 string) (*MatchPrediction, error)
	GetSeasonOutlook() (*SeasonSimulation, error)
}

type AdvancedPredictionService struct {
	pythonExecutable string
	scriptPath       string
}

// NewAdvancedPredictionService creates a new prediction service
func NewAdvancedPredictionService() *AdvancedPredictionService {
	return &AdvancedPredictionService{
		pythonExecutable: detectPythonCommand(),
		scriptPath:       "prediction.py",
	}
}

// detectPythonCommand detects the available Python command
func detectPythonCommand() string {
	commands := []string{"python3", "python"}

	for _, cmd := range commands {
		_, err := exec.LookPath(cmd)
		if err == nil {
			return cmd
		}
	}

	// Default fallback based on OS
	if runtime.GOOS == "windows" {
		return "python"
	}
	return "python3"
}

// prediction.go
func (aps *AdvancedPredictionService) RunAdvancedPrediction() (*ComprehensivePrediction, error) {
	// Get historical matches from database
	historicalMatches, err := db.GetHistoricalMatches()
	if err != nil {
		log.Printf("Warning: could not fetch historical matches: %v", err)
	}

	// Convert to JSON to pass to Python script
	historicalData, err := json.Marshal(historicalMatches)
	if err != nil {
		log.Printf("Warning: could not marshal historical data: %v", err)
	}

	cmd := exec.Command(aps.pythonExecutable, aps.scriptPath)

	// Create stdin pipe
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to create stdin pipe: %v", err)
	}

	// Write historical data to stdin
	go func() {
		defer stdin.Close()
		stdin.Write(historicalData)
	}()

	out, err := cmd.Output()
	if err != nil {
		// Fallback to alternative Python command if needed
		altPython := "python"
		if aps.pythonExecutable == "python" {
			altPython = "python3"
		}
		cmd = exec.Command(altPython, aps.scriptPath)

		// Try again with historical data
		stdin, err = cmd.StdinPipe()
		if err == nil {
			go func() {
				defer stdin.Close()
				stdin.Write(historicalData)
			}()
			out, err = cmd.Output()
		}

		if err != nil {
			return nil, fmt.Errorf("failed to execute prediction script: %v", err)
		}
	}

	var prediction ComprehensivePrediction
	err = json.Unmarshal(out, &prediction)
	if err != nil {
		return nil, fmt.Errorf("failed to parse prediction results: %v", err)
	}

	return &prediction, nil
}

// GetMatchPrediction gets prediction for a specific match
func (aps *AdvancedPredictionService) GetMatchPrediction(team1, team2 string) (*MatchPrediction, error) {
	prediction, err := aps.RunAdvancedPrediction()
	if err != nil {
		return nil, err
	}

	// Find the specific match prediction
	for _, match := range prediction.MatchPredictions {
		if (match.Team1 == team1 && match.Team2 == team2) ||
			(match.Team1 == team2 && match.Team2 == team1) {
			return &match, nil
		}
	}

	return nil, fmt.Errorf("match prediction not found for %s vs %s", team1, team2)
}

// GetSeasonOutlook gets the season simulation results
func (aps *AdvancedPredictionService) GetSeasonOutlook() (*SeasonSimulation, error) {
	prediction, err := aps.RunAdvancedPrediction()
	if err != nil {
		return nil, err
	}

	return &prediction.SeasonSimulation, nil
}

// Legacy function for backward compatibility
func RunPrediction() ([]MatchPrediction, error) {
	service := NewAdvancedPredictionService()
	prediction, err := service.RunAdvancedPrediction()
	if err != nil {
		return nil, err
	}

	return prediction.MatchPredictions, nil
}

// PredictionAnalytics provides analytics on prediction accuracy
type PredictionAnalytics struct {
	TotalPredictions   int     `json:"total_predictions"`
	CorrectPredictions int     `json:"correct_predictions"`
	AccuracyPercentage float64 `json:"accuracy_percentage"`
	AverageConfidence  float64 `json:"average_confidence"`
	LastCalculated     string  `json:"last_calculated"`
}

// CalculateAccuracy calculates prediction accuracy (for future implementation)
func (aps *AdvancedPredictionService) CalculateAccuracy() *PredictionAnalytics {
	// This would be implemented with actual match results
	// For now, return placeholder data
	return &PredictionAnalytics{
		TotalPredictions:   100,
		CorrectPredictions: 73,
		AccuracyPercentage: 73.0,
		AverageConfidence:  82.5,
		LastCalculated:     time.Now().Format(time.RFC3339),
	}
}
