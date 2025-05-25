package db

import "leaguesimulator/models"

func SaveMatch(match models.Match) error {
	query := `
		INSERT INTO matches (week, home_team_name, away_team_name, home_goals, away_goals, played)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(query,
		match.Week,
		match.HomeTeam,
		match.AwayTeam,
		match.HomeGoals,
		match.AwayGoals,
		match.Played,
	)

	return err
}
