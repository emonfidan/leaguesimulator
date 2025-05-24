package league

import (
	"math/rand"
	"sort"
	"time"

	"leaguesimulator/models"
)

type LeagueManager struct {
	Teams     []models.Team
	Matches   []models.Match
	Week      int
	Standings []TeamStanding
}

type Match struct {
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

// Initialize league
func (lm *LeagueManager) InitLeague() {
	lm.Teams = []models.Team{
		{Name: "Lions", Strength: 90},
		{Name: "Tigers", Strength: 80},
		{Name: "Bears", Strength: 70},
		{Name: "Wolves", Strength: 60},
	}
	lm.Week = 0
	lm.Matches = []models.Match{}
	lm.Standings = []TeamStanding{}
}

// Play match between two teams
func (lm *LeagueManager) playMatch(week int, home *models.Team, away *models.Team) models.Match {
	rand.Seed(time.Now().UnixNano())
	// Generate goals based on strength
	homeGoals := rand.Intn(home.Strength/15 + 1)
	awayGoals := rand.Intn(away.Strength/15 + 1)

	// Update team stats
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

	return models.Match{
		Week:      week,
		HomeTeam:  home.Name,
		AwayTeam:  away.Name,
		HomeGoals: homeGoals,
		AwayGoals: awayGoals,
		Played:    true,
		Team1:     home.Name,
		Team2:     away.Name,
		Score1:    homeGoals,
		Score2:    awayGoals,
	}
}

// Play next week matches
func (lm *LeagueManager) PlayNextWeek() []Match {
	if lm.Week >= 3 { // 4 teams = 3 rounds in round-robin (each team plays each other once)
		return nil
	}

	// Define matchups for each week
	matchups := [][][]int{
		{{0, 1}, {2, 3}}, // Week 1: Lions vs Tigers, Bears vs Wolves
		{{0, 2}, {1, 3}}, // Week 2: Lions vs Bears, Tigers vs Wolves
		{{0, 3}, {1, 2}}, // Week 3: Lions vs Wolves, Tigers vs Bears
	}

	if lm.Week < len(matchups) {
		weekMatchups := matchups[lm.Week]
		var weekMatches []Match

		for _, matchup := range weekMatchups {
			match := lm.playMatch(lm.Week+1, &lm.Teams[matchup[0]], &lm.Teams[matchup[1]])
			lm.Matches = append(lm.Matches, match)

			weekMatches = append(weekMatches, Match{
				Week:   match.Week,
				Team1:  match.HomeTeam,
				Team2:  match.AwayTeam,
				Score1: match.HomeGoals,
				Score2: match.AwayGoals,
			})
		}

		lm.Week++
		lm.updateStandings()
		return weekMatches
	}

	return nil
}

// Update standings based on current matches
func (lm *LeagueManager) updateStandings() {
	standings := make(map[string]*TeamStanding)

	// Initialize standings
	for _, team := range lm.Teams {
		standings[team.Name] = &TeamStanding{
			Name: team.Name,
		}
	}

	// Calculate standings from matches
	for _, match := range lm.Matches {
		home := standings[match.HomeTeam]
		away := standings[match.AwayTeam]

		home.Played++
		away.Played++

		home.GoalsFor += match.HomeGoals
		home.GoalsAgainst += match.AwayGoals
		away.GoalsFor += match.AwayGoals
		away.GoalsAgainst += match.HomeGoals

		if match.HomeGoals > match.AwayGoals {
			home.Won++
			away.Lost++
			home.Points += 3
		} else if match.HomeGoals < match.AwayGoals {
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

	// Calculate goal difference
	for _, standing := range standings {
		standing.GoalDiff = standing.GoalsFor - standing.GoalsAgainst
	}

	// Convert to slice and sort
	var sortedStandings []TeamStanding
	for _, s := range standings {
		sortedStandings = append(sortedStandings, *s)
	}

	// Sort by points, then goal difference, then goals for
	sort.Slice(sortedStandings, func(i, j int) bool {
		if sortedStandings[i].Points != sortedStandings[j].Points {
			return sortedStandings[i].Points > sortedStandings[j].Points
		}
		if sortedStandings[i].GoalDiff != sortedStandings[j].GoalDiff {
			return sortedStandings[i].GoalDiff > sortedStandings[j].GoalDiff
		}
		return sortedStandings[i].GoalsFor > sortedStandings[j].GoalsFor
	})

	lm.Standings = sortedStandings
}

// UpdateStandings - public method for router
func (lm *LeagueManager) UpdateStandings() {
	lm.updateStandings()
}

// Get current standings
func (lm *LeagueManager) GetStandings() []TeamStanding {
	lm.updateStandings()
	return lm.Standings
}

// Get all matches played
func (lm *LeagueManager) GetMatches() []models.Match {
	return lm.Matches
}

// Get future fixtures
func (lm *LeagueManager) GetFutureFixtures() []Match {
	var fixtures []Match

	// Generate all possible matchups for remaining weeks
	allMatchups := [][][]string{
		{{"Lions", "Tigers"}, {"Bears", "Wolves"}}, // Week 1
		{{"Lions", "Bears"}, {"Tigers", "Wolves"}}, // Week 2
		{{"Lions", "Wolves"}, {"Tigers", "Bears"}}, // Week 3
	}

	// Check which weeks haven't been played yet
	for week := lm.Week; week < 3; week++ {
		weekMatchups := allMatchups[week]
		for _, matchup := range weekMatchups {
			fixtures = append(fixtures, Match{
				Week:  week + 1,
				Team1: matchup[0],
				Team2: matchup[1],
			})
		}
	}

	return fixtures
}

// Edit match result
func (lm *LeagueManager) EditMatchResult(week int, team1, team2 string, score1, score2 int) bool {
	for i, match := range lm.Matches {
		if match.Week == week &&
			((match.HomeTeam == team1 && match.AwayTeam == team2) ||
				(match.HomeTeam == team2 && match.AwayTeam == team1)) {
			// Update the match
			if match.HomeTeam == team1 {
				lm.Matches[i].HomeGoals = score1
				lm.Matches[i].AwayGoals = score2
			} else {
				lm.Matches[i].HomeGoals = score2
				lm.Matches[i].AwayGoals = score1
			}

			// Update compatibility fields
			lm.Matches[i].Score1 = score1
			lm.Matches[i].Score2 = score2

			// Recalculate team stats
			lm.recalculateTeamStats()
			return true
		}
	}
	return false
}

// Recalculate team statistics after editing match results
func (lm *LeagueManager) recalculateTeamStats() {
	// Reset all team stats
	for i := range lm.Teams {
		lm.Teams[i].Points = 0
		lm.Teams[i].Played = 0
		lm.Teams[i].Wins = 0
		lm.Teams[i].Draws = 0
		lm.Teams[i].Losses = 0
		lm.Teams[i].GoalsFor = 0
		lm.Teams[i].GoalsAgainst = 0
	}

	// Recalculate from all matches
	teamMap := make(map[string]*models.Team)
	for i := range lm.Teams {
		teamMap[lm.Teams[i].Name] = &lm.Teams[i]
	}

	for _, match := range lm.Matches {
		home := teamMap[match.HomeTeam]
		away := teamMap[match.AwayTeam]

		home.Played++
		away.Played++
		home.GoalsFor += match.HomeGoals
		home.GoalsAgainst += match.AwayGoals
		away.GoalsFor += match.AwayGoals
		away.GoalsAgainst += match.HomeGoals

		if match.HomeGoals > match.AwayGoals {
			home.Points += 3
			home.Wins++
			away.Losses++
		} else if match.AwayGoals > match.HomeGoals {
			away.Points += 3
			away.Wins++
			home.Losses++
		} else {
			home.Points++
			away.Points++
			home.Draws++
			away.Draws++
		}
	}

	lm.updateStandings()
}
