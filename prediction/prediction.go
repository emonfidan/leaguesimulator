package prediction

import (
	"encoding/json"
	"os/exec"
)

type MatchPrediction struct {
	Team1  string `json:"team1"`
	Team2  string `json:"team2"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
	Result string `json:"result"`
}

func RunPrediction() ([]MatchPrediction, error) {
	cmd := exec.Command("python", "prediction.py")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var predictions []MatchPrediction
	err = json.Unmarshal(out, &predictions)
	if err != nil {
		return nil, err
	}

	return predictions, nil
}
