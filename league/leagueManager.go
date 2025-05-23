package league

import (
	"math/rand"
	"time"

	"leaguesimulator/models"
)

type LeagueManager struct {
	Teams   []models.Team
	Matches []models.Match
	Week    int
}

// Lig başlatma
func (lm *LeagueManager) InitLeague() {
	lm.Teams = []models.Team{
		{Name: "Lions", Strength: 90},
		{Name: "Tigers", Strength: 80},
		{Name: "Bears", Strength: 70},
		{Name: "Wolves", Strength: 60},
	}
	lm.Week = 0
	lm.Matches = []models.Match{}
}

// Belirli iki takımın maçını oynat
func (lm *LeagueManager) playMatch(week int, home *models.Team, away *models.Team) models.Match {
	rand.Seed(time.Now().UnixNano())
	// Güce göre rastgele gol üret
	homeGoals := rand.Intn(home.Strength/10 + 1)
	awayGoals := rand.Intn(away.Strength/10 + 1)

	// Puanlama
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
	}
}

// Haftayı oynat
func (lm *LeagueManager) PlayNextWeek() []models.Match {
	if lm.Week >= 6 { // 4 takımlı ligde toplam 6 hafta olur
		return nil
	}

	matchups := [][]int{
		{0, 1}, {2, 3}, {0, 2}, {1, 3}, {0, 3}, {1, 2},
	}

	m := matchups[lm.Week]
	match1 := lm.playMatch(lm.Week+1, &lm.Teams[m[0]], &lm.Teams[m[1]])

	lm.Matches = append(lm.Matches, match1)
	lm.Week++

	return []models.Match{match1}
}

// Lig tablosunu döndür
func (lm *LeagueManager) GetStandings() []models.Team {
	return lm.Teams
}

// Oynanan maçları döndür
func (lm *LeagueManager) GetMatches() []models.Match {
	return lm.Matches
}
