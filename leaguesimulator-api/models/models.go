package models

type Team struct {
	Name         string `json:"name"`
	Points       int    `json:"points"`
	Played       int    `json:"played"`
	Wins         int    `json:"wins"`
	Draws        int    `json:"draws"`
	Losses       int    `json:"losses"`
	GoalsFor     int    `json:"goals_for"`
	GoalsAgainst int    `json:"goals_against"`
	Strength     int    `json:"strength"`
}

type Match struct {
	ID        int    `json:"id"`
	Week      int    `json:"week"`
	HomeTeam  string `json:"home_team"`
	AwayTeam  string `json:"away_team"`
	HomeGoals int    `json:"home_goals"`
	AwayGoals int    `json:"away_goals"`
	Played    bool   `json:"played"`
}

type Prediction struct {
	ID            int    `json:"id"`
	TeamName      string `json:"team_name"`
	PredictedRank int    `json:"predicted_rank"`
	WeekSubmitted int    `json:"week_submitted"`
}

type HistoricalMatch struct {
	ID        int    `json:"id"`
	Season    int    `json:"season"`
	Week      int    `json:"week"`
	HomeTeam  string `json:"home_team"`
	AwayTeam  string `json:"away_team"`
	HomeGoals int    `json:"home_goals"`
	AwayGoals int    `json:"away_goals"`
}
