export interface Team {
  name: string;
  played: number;
  won: number;
  drawn: number;
  lost: number;
  goals_for: number;
  goals_against: number;
  goal_diff: number;
  points: number;
}

export interface Match {
  week: number;
  team1: string;
  team2: string;
  score1: number;
  score2: number;
  result?: string;
}

export interface Prediction {
  team1: string;
  team2: string;
  score1: number;
  score2: number;
  result: string;
  confidence: number;
  weather: string;
  expected_goals: {
    home: number;
    away: number;
  };
  win_probabilities: {
    home_win: number;
    draw: number;
    away_win: number;
  };
}

export interface SeasonSimulation {
  championship_probabilities: Record<string, number>;
  simulation_runs: number;
}

export interface StandingsResponse {
  current_week: number;
  standings: Team[];
  total_teams: number;
  league_status: string;
}

export interface TeamAnalysis {
  team: string;
  performance: {
    home: {
      wins: number;
      draws: number;
      losses: number;
    };
    away: {
      wins: number;
      draws: number;
      losses: number;
    };
  };
  form: string[];
  goals: {
    for: {
      average: number;
      distribution: Record<string, number>;
    };
    against: {
      average: number;
      distribution: Record<string, number>;
    };
  };
}