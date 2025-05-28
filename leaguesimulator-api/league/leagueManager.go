package league

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"

	"leaguesimulator/db"
	"leaguesimulator/models"
)

type LeagueManager struct {
	Teams     []models.Team
	Matches   []models.Match
	Week      int
	Standings []TeamStanding
}

type MatchView struct {
	Week   int    `json:"week"`
	Team1  string `json:"team1"`
	Team2  string `json:"team2"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
}

type TeamStanding struct {
	Name         string `json:"name"`
	Played       int    `json:"played"`
	Won          int    `json:"won"`
	Drawn        int    `json:"drawn"`
	Lost         int    `json:"lost"`
	GoalsFor     int    `json:"goals_for"`
	GoalsAgainst int    `json:"goals_against"`
	GoalDiff     int    `json:"goal_diff"`
	Points       int    `json:"points"`
}

// InitLeague initializes teams and resets stats
func (lm *LeagueManager) InitLeague() {
	teams, err := db.GetAllTeams()
	if err != nil || len(teams) == 0 {
		// Create default teams
		teams = []models.Team{
			{Name: "Lions", Strength: 90},
			{Name: "Tigers", Strength: 80},
			{Name: "Bears", Strength: 70},
			{Name: "Wolves", Strength: 60},
		}
		_ = db.SaveTeams(teams)
	}

	lm.Teams = teams
	lm.Week = 0

	// Load existing matches from database
	matches, err := db.GetAllMatches()
	if err == nil {
		lm.Matches = matches
		// Calculate current week from existing matches
		maxWeek := 0
		for _, match := range matches {
			if match.Week > maxWeek {
				maxWeek = match.Week
			}
		}
		lm.Week = maxWeek
	} else {
		lm.Matches = []models.Match{}
	}

	lm.Standings = []TeamStanding{}
	lm.updateStandings()
}

// playMatch simulates a match between home and away teams, updates their stats, returns the match record
func (lm *LeagueManager) playMatch(week int, home *models.Team, away *models.Team) models.Match {
	rand.Seed(time.Now().UnixNano())

	homeGoals := rand.Intn(home.Strength/15 + 1)
	awayGoals := rand.Intn(away.Strength/15 + 1)

	if homeGoals > awayGoals {
		home.Points += 3
		home.Wins++
		away.Losses++
	} else if awayGoals > homeGoals {
		away.Points += 3
		away.Wins++
		home.Losses++
	} else {
		home.Points++
		away.Points++
		home.Draws++
		away.Draws++
	}

	home.Played++
	away.Played++
	home.GoalsFor += homeGoals
	home.GoalsAgainst += awayGoals
	away.GoalsFor += awayGoals
	away.GoalsAgainst += homeGoals

	// Save updated team stats to database
	_ = db.SaveSingleTeam(*home)
	_ = db.SaveSingleTeam(*away)

	match := models.Match{
		Week:      week,
		HomeTeam:  home.Name,
		AwayTeam:  away.Name,
		HomeGoals: homeGoals,
		AwayGoals: awayGoals,
		Played:    true,
	}

	// Save to historical matches
	err := db.SaveHistoricalMatch(match)
	if err != nil {
		log.Printf("Failed to save historical match: %v", err)
	}

	return match
}

// PlayNextWeek simulates the matches of the next week, updates matches and standings
func (lm *LeagueManager) PlayNextWeek() []MatchView {
	if lm.Week >= 6 { // 6 rounds for 4 teams double round robin
		return nil
	}

	matchups := [][][]int{
		{{0, 1}, {2, 3}}, // Week 1
		{{0, 2}, {1, 3}}, // Week 2
		{{0, 3}, {1, 2}}, // Week 3
		{{1, 0}, {3, 2}}, // Week 4 (reverse)
		{{2, 0}, {3, 1}}, // Week 5
		{{3, 0}, {2, 1}}, // Week 6
	}

	weekMatchups := matchups[lm.Week]
	var playedMatches []MatchView

	for _, m := range weekMatchups {
		match := lm.playMatch(lm.Week+1, &lm.Teams[m[0]], &lm.Teams[m[1]])
		_ = db.SaveMatch(match)
		lm.Matches = append(lm.Matches, match)

		playedMatches = append(playedMatches, MatchView{
			Week:   match.Week,
			Team1:  match.HomeTeam,
			Team2:  match.AwayTeam,
			Score1: match.HomeGoals,
			Score2: match.AwayGoals,
		})
	}

	lm.Week++
	lm.updateStandings()
	return playedMatches
}

// updateStandings recalculates the league table from matches
func (lm *LeagueManager) updateStandings() {
	standings := make(map[string]*TeamStanding)

	for _, t := range lm.Teams {
		standings[t.Name] = &TeamStanding{Name: t.Name}
	}

	for _, m := range lm.Matches {
		home := standings[m.HomeTeam]
		away := standings[m.AwayTeam]

		home.Played++
		away.Played++

		home.GoalsFor += m.HomeGoals
		home.GoalsAgainst += m.AwayGoals
		away.GoalsFor += m.AwayGoals
		away.GoalsAgainst += m.HomeGoals

		if m.HomeGoals > m.AwayGoals {
			home.Won++
			away.Lost++
			home.Points += 3
		} else if m.AwayGoals > m.HomeGoals {
			away.Won++
			home.Lost++
			away.Points += 3
		} else {
			home.Drawn++
			away.Drawn++
			home.Points++
			away.Points++
		}
	}

	for _, s := range standings {
		s.GoalDiff = s.GoalsFor - s.GoalsAgainst
	}

	// Convert to slice
	lm.Standings = []TeamStanding{}
	for _, s := range standings {
		lm.Standings = append(lm.Standings, *s)
	}

	sort.Slice(lm.Standings, func(i, j int) bool {
		if lm.Standings[i].Points != lm.Standings[j].Points {
			return lm.Standings[i].Points > lm.Standings[j].Points
		}
		if lm.Standings[i].GoalDiff != lm.Standings[j].GoalDiff {
			return lm.Standings[i].GoalDiff > lm.Standings[j].GoalDiff
		}
		return lm.Standings[i].GoalsFor > lm.Standings[j].GoalsFor
	})
}

// UpdateStandings is a public method for external calls
func (lm *LeagueManager) UpdateStandings() {
	lm.updateStandings()
}

// GetStandings returns current standings
func (lm *LeagueManager) GetStandings() []TeamStanding {
	lm.updateStandings()
	return lm.Standings
}

// GetMatches returns all played matches
func (lm *LeagueManager) GetMatches() []models.Match {
	// Reload matches from database to ensure consistency
	matches, err := db.GetAllMatches()
	if err == nil {
		lm.Matches = matches
	}
	return lm.Matches
}

// GetFutureFixtures returns matches not yet played (simple hardcoded fixture list)
func (lm *LeagueManager) GetFutureFixtures() []MatchView {
	var fixtures []MatchView
	allMatchups := [][][]string{
		{{"Lions", "Tigers"}, {"Bears", "Wolves"}},
		{{"Lions", "Bears"}, {"Tigers", "Wolves"}},
		{{"Lions", "Wolves"}, {"Tigers", "Bears"}},
		{{"Tigers", "Lions"}, {"Wolves", "Bears"}},
		{{"Bears", "Lions"}, {"Wolves", "Tigers"}},
		{{"Wolves", "Lions"}, {"Bears", "Tigers"}},
	}

	for w := lm.Week; w < len(allMatchups); w++ {
		for _, m := range allMatchups[w] {
			fixtures = append(fixtures, MatchView{
				Week:  w + 1,
				Team1: m[0],
				Team2: m[1],
			})
		}
	}

	return fixtures
}

// EditMatchResult edits a played match and recalculates stats
func (lm *LeagueManager) EditMatchResult(week int, team1, team2 string, score1, score2 int) bool {
	for i, m := range lm.Matches {
		if m.Week == week &&
			((m.HomeTeam == team1 && m.AwayTeam == team2) || (m.HomeTeam == team2 && m.AwayTeam == team1)) {
			if m.HomeTeam == team1 {
				lm.Matches[i].HomeGoals = score1
				lm.Matches[i].AwayGoals = score2
			} else {
				lm.Matches[i].HomeGoals = score2
				lm.Matches[i].AwayGoals = score1
			}

			// Update in database
			_ = db.UpdateMatch(lm.Matches[i])

			lm.recalculateTeamStats()
			return true
		}
	}
	return false
}

// recalculateTeamStats resets and recalculates all teams stats from played matches
func (lm *LeagueManager) recalculateTeamStats() {
	// Reset all team stats in database
	_ = db.ResetAllTeamStats()

	// Reset local team stats
	for i := range lm.Teams {
		lm.Teams[i].Played = 0
		lm.Teams[i].Wins = 0
		lm.Teams[i].Draws = 0
		lm.Teams[i].Losses = 0
		lm.Teams[i].GoalsFor = 0
		lm.Teams[i].GoalsAgainst = 0
		lm.Teams[i].Points = 0
	}

	for _, m := range lm.Matches {
		var home, away *models.Team
		for i := range lm.Teams {
			if lm.Teams[i].Name == m.HomeTeam {
				home = &lm.Teams[i]
			}
			if lm.Teams[i].Name == m.AwayTeam {
				away = &lm.Teams[i]
			}
		}

		if home == nil || away == nil {
			continue
		}

		home.Played++
		away.Played++
		home.GoalsFor += m.HomeGoals
		home.GoalsAgainst += m.AwayGoals
		away.GoalsFor += m.AwayGoals
		away.GoalsAgainst += m.HomeGoals

		if m.HomeGoals > m.AwayGoals {
			home.Wins++
			away.Losses++
			home.Points += 3
		} else if m.AwayGoals > m.HomeGoals {
			away.Wins++
			home.Losses++
			away.Points += 3
		} else {
			home.Draws++
			away.Draws++
			home.Points++
			away.Points++
		}
	}

	// Save updated team stats to database
	_ = db.SaveTeams(lm.Teams)

	lm.updateStandings()
}

// ResetLeague clears all matches, resets weeks and team stats
func (lm *LeagueManager) ResetLeague() {
	// Clear matches from database
	_ = db.ClearAllMatches()

	// Reset team stats in database
	_ = db.ResetAllTeamStats()

	// Reset local data
	lm.Matches = []models.Match{}
	lm.Week = 0
	for i := range lm.Teams {
		lm.Teams[i].Played = 0
		lm.Teams[i].Wins = 0
		lm.Teams[i].Draws = 0
		lm.Teams[i].Losses = 0
		lm.Teams[i].GoalsFor = 0
		lm.Teams[i].GoalsAgainst = 0
		lm.Teams[i].Points = 0
	}
	lm.updateStandings()
}

// GetMatchById returns a match by its ID (index in matches slice)
func (lm *LeagueManager) GetMatchById(matchId int) (models.Match, error) {
	// Reload matches from database to ensure consistency
	matches, err := db.GetAllMatches()
	if err == nil {
		lm.Matches = matches
	}

	// Check if matchId is valid (0-based index)
	if matchId < 0 || matchId >= len(lm.Matches) {
		return models.Match{}, fmt.Errorf("match with ID %d not found", matchId)
	}

	return lm.Matches[matchId], nil
}

// EditMatchResultById edits a match by its ID and recalculates stats
func (lm *LeagueManager) EditMatchResultById(matchId int, homeGoals, awayGoals int) bool {
	// Reload matches from database to ensure consistency
	matches, err := db.GetAllMatches()
	if err == nil {
		lm.Matches = matches
	}

	// Check if matchId is valid
	if matchId < 0 || matchId >= len(lm.Matches) {
		return false
	}

	// Validate goals (should be non-negative)
	if homeGoals < 0 || awayGoals < 0 {
		return false
	}

	// Update the match result
	lm.Matches[matchId].HomeGoals = homeGoals
	lm.Matches[matchId].AwayGoals = awayGoals

	// Update in database
	err = db.UpdateMatch(lm.Matches[matchId])
	if err != nil {
		return false
	}

	// Recalculate all team stats based on updated matches
	lm.recalculateTeamStats()

	return true
}
