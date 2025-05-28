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

func GetAllMatches() ([]models.Match, error) {
	query := `SELECT id, week, home_team_name, away_team_name, home_goals, away_goals, played FROM matches ORDER BY week, id`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []models.Match
	for rows.Next() {
		var match models.Match
		err := rows.Scan(
			&match.ID,
			&match.Week,
			&match.HomeTeam,
			&match.AwayTeam,
			&match.HomeGoals,
			&match.AwayGoals,
			&match.Played,
		)
		if err != nil {
			return nil, err
		}
		matches = append(matches, match)
	}
	return matches, nil
}

func UpdateMatch(match models.Match) error {
	query := `
		UPDATE matches 
		SET home_goals = ?, away_goals = ?, played = ?
		WHERE week = ? AND home_team_name = ? AND away_team_name = ?
	`

	_, err := DB.Exec(query,
		match.HomeGoals,
		match.AwayGoals,
		match.Played,
		match.Week,
		match.HomeTeam,
		match.AwayTeam,
	)

	return err
}

func ClearAllMatches() error {
	query := `DELETE FROM matches`
	_, err := DB.Exec(query)
	return err
}

// match_repository.go
func SaveHistoricalMatch(match models.Match) error {
	query := `
        INSERT INTO historical_matches 
        (season, week, home_team_name, away_team_name, home_goals, away_goals)
        VALUES (1, ?, ?, ?, ?, ?)
    `
	_, err := DB.Exec(query,
		match.Week,
		match.HomeTeam,
		match.AwayTeam,
		match.HomeGoals,
		match.AwayGoals,
	)
	return err
}
