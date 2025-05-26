package db

import (
	"leaguesimulator/models"
)

func GetAllTeams() ([]models.Team, error) {
	query := `SELECT name, points, played, wins, draws, losses, goals_for, goals_against, strength FROM teams ORDER BY name`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err := rows.Scan(
			&team.Name,
			&team.Points,
			&team.Played,
			&team.Wins,
			&team.Draws,
			&team.Losses,
			&team.GoalsFor,
			&team.GoalsAgainst,
			&team.Strength,
		)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	return teams, nil
}

func SaveTeams(teams []models.Team) error {
	query := `
		INSERT INTO teams 
		(name, points, played, wins, draws, losses, goals_for, goals_against, strength)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE 
		points=VALUES(points),
		played=VALUES(played),
		wins=VALUES(wins),
		draws=VALUES(draws),
		losses=VALUES(losses),
		goals_for=VALUES(goals_for),
		goals_against=VALUES(goals_against),
		strength=VALUES(strength)
	`

	for _, team := range teams {
		_, err := DB.Exec(query,
			team.Name,
			team.Points,
			team.Played,
			team.Wins,
			team.Draws,
			team.Losses,
			team.GoalsFor,
			team.GoalsAgainst,
			team.Strength,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveSingleTeam(team models.Team) error {
	query := `
		INSERT INTO teams 
		(name, points, played, wins, draws, losses, goals_for, goals_against, strength)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE 
		points=VALUES(points),
		played=VALUES(played),
		wins=VALUES(wins),
		draws=VALUES(draws),
		losses=VALUES(losses),
		goals_for=VALUES(goals_for),
		goals_against=VALUES(goals_against),
		strength=VALUES(strength)
	`

	_, err := DB.Exec(query,
		team.Name,
		team.Points,
		team.Played,
		team.Wins,
		team.Draws,
		team.Losses,
		team.GoalsFor,
		team.GoalsAgainst,
		team.Strength,
	)

	return err
}

func ResetTeamStats(teamName string) error {
	query := `
		UPDATE teams 
		SET points = 0, played = 0, wins = 0, draws = 0, losses = 0, 
		    goals_for = 0, goals_against = 0
		WHERE name = ?
	`
	_, err := DB.Exec(query, teamName)
	return err
}

func ResetAllTeamStats() error {
	query := `
		UPDATE teams 
		SET points = 0, played = 0, wins = 0, draws = 0, losses = 0, 
		    goals_for = 0, goals_against = 0
	`
	_, err := DB.Exec(query)
	return err
}
