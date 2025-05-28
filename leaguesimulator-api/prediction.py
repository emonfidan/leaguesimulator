import random
import json
import math
import sys
from datetime import datetime

class AdvancedPredictor:
    def __init__(self):
        # Enhanced team data with multiple attributes
        self.historical_data = {}
        self.current_season_results = {}  # Track current season match results
        self.teams_data = {
            "Lions": {
                "attack": 90,
                "defense": 85,
                "midfield": 88,
                "form": 0.8,  # Recent form factor
                "home_advantage": 1.15,
                "fatigue": 0.0,
                "injury_rate": 0.05,
                "historical_performance": [],
                "season_momentum": 1.0,  # New: momentum based on current season
                "confidence_boost": 1.0   # New: confidence from early wins
            },
            "Tigers": {
                "attack": 82,
                "defense": 78,
                "midfield": 80,
                "form": 0.75,
                "home_advantage": 1.12,
                "fatigue": 0.0,
                "injury_rate": 0.08,
                "historical_performance": [],
                "season_momentum": 1.0,
                "confidence_boost": 1.0
            },
            "Bears": {
                "attack": 75,
                "defense": 80,
                "midfield": 72,
                "form": 0.7,
                "home_advantage": 1.1,
                "fatigue": 0.0,
                "injury_rate": 0.1,
                "historical_performance": [],
                "season_momentum": 1.0,
                "confidence_boost": 1.0
            },
            "Wolves": {
                "attack": 68,
                "defense": 70,
                "midfield": 65,
                "form": 0.65,
                "home_advantage": 1.08,
                "fatigue": 0.0,
                "injury_rate": 0.12,
                "historical_performance": [],
                "season_momentum": 1.0,
                "confidence_boost": 1.0
            }
        }
        
        # Weather impact on performance
        self.weather_conditions = ["sunny", "rainy", "windy", "cloudy"]
        self.weather_impact = {
            "sunny": 1.0,
            "rainy": 0.9,
            "windy": 0.85,
            "cloudy": 0.95
        }
        
        # Match importance factor - Enhanced for early season matches
        self.match_importance = {
            "early_season": 1.4,  # First 4 matches have higher importance
            "regular": 1.0,
            "derby": 1.2,
            "final_week": 1.3
        }

    def load_historical_data(self, historical_matches):
        """Process historical match data from database"""
        for match in historical_matches:
            home_team = match['home_team']
            away_team = match['away_team']
            key = f"{home_team}_{away_team}"
            
            if key not in self.historical_data:
                self.historical_data[key] = []
            
            self.historical_data[key].append({
                'home_goals': match['home_goals'],
                'away_goals': match['away_goals'],
                'result': self._get_result(match['home_goals'], match['away_goals']),
                'season': match.get('season', 1),
                'week': match['week']
            })
            
            # Update team historical performance
            for team in [home_team, away_team]:
                if team in self.teams_data:
                    self.teams_data[team]['historical_performance'].append({
                        'goals_for': match['home_goals'] if team == home_team else match['away_goals'],
                        'goals_against': match['away_goals'] if team == home_team else match['home_goals'],
                        'result': 'win' if (team == home_team and match['home_goals'] > match['away_goals']) or 
                                      (team == away_team and match['away_goals'] > match['home_goals']) else 
                                 'draw' if match['home_goals'] == match['away_goals'] else 'loss'
                    })

    def update_current_season_results(self, match_result):
        """Update current season results to influence future predictions"""
        home_team = match_result['home_team']
        away_team = match_result['away_team']
        home_goals = match_result['home_goals']
        away_goals = match_result['away_goals']
        week = match_result['week']
        
        # Store match result
        if week not in self.current_season_results:
            self.current_season_results[week] = []
        
        self.current_season_results[week].append({
            'home_team': home_team,
            'away_team': away_team,
            'home_goals': home_goals,
            'away_goals': away_goals,
            'result': self._get_result(home_goals, away_goals)
        })
        
        # Update team momentum and confidence based on results
        self._update_team_momentum(home_team, away_team, home_goals, away_goals, week)

    def _update_team_momentum(self, home_team, away_team, home_goals, away_goals, week):
        """Update team momentum based on match results - higher impact for early matches"""
        # Calculate impact multiplier (higher for first 4 matches)
        impact_multiplier = 2.0 if week <= 4 else 1.0
        
        # Determine result and update momentum
        if home_goals > away_goals:
            # Home team wins
            self.teams_data[home_team]['season_momentum'] *= (1.0 + 0.15 * impact_multiplier)
            self.teams_data[home_team]['confidence_boost'] *= (1.0 + 0.1 * impact_multiplier)
            self.teams_data[away_team]['season_momentum'] *= (1.0 - 0.1 * impact_multiplier)
            self.teams_data[away_team]['confidence_boost'] *= (1.0 - 0.05 * impact_multiplier)
        elif away_goals > home_goals:
            # Away team wins
            self.teams_data[away_team]['season_momentum'] *= (1.0 + 0.18 * impact_multiplier)  # Slightly higher for away wins
            self.teams_data[away_team]['confidence_boost'] *= (1.0 + 0.12 * impact_multiplier)
            self.teams_data[home_team]['season_momentum'] *= (1.0 - 0.1 * impact_multiplier)
            self.teams_data[home_team]['confidence_boost'] *= (1.0 - 0.05 * impact_multiplier)
        else:
            # Draw - smaller impact
            self.teams_data[home_team]['season_momentum'] *= (1.0 + 0.02 * impact_multiplier)
            self.teams_data[away_team]['season_momentum'] *= (1.0 + 0.02 * impact_multiplier)
        
        # Apply bounds to prevent extreme values
        for team in [home_team, away_team]:
            self.teams_data[team]['season_momentum'] = max(0.5, min(2.0, self.teams_data[team]['season_momentum']))
            self.teams_data[team]['confidence_boost'] = max(0.7, min(1.5, self.teams_data[team]['confidence_boost']))

    def calculate_current_season_form(self, team_name, current_week):
        """Calculate form based on current season results with heavy weighting on first 4 matches"""
        team_results = []
        
        # Collect all results for this team up to current week
        for week in range(1, current_week):
            if week in self.current_season_results:
                for match in self.current_season_results[week]:
                    if match['home_team'] == team_name:
                        if match['home_goals'] > match['away_goals']:
                            team_results.append(('win', week))
                        elif match['home_goals'] < match['away_goals']:
                            team_results.append(('loss', week))
                        else:
                            team_results.append(('draw', week))
                    elif match['away_team'] == team_name:
                        if match['away_goals'] > match['home_goals']:
                            team_results.append(('win', week))
                        elif match['away_goals'] < match['home_goals']:
                            team_results.append(('loss', week))
                        else:
                            team_results.append(('draw', week))
        
        if not team_results:
            return 1.0  # Neutral form if no matches played
        
        # Calculate weighted form - first 4 matches have 3x weight
        total_weight = 0
        weighted_score = 0
        
        for result, week in team_results:
            weight = 3.0 if week <= 4 else 1.0  # Triple weight for first 4 matches
            total_weight += weight
            
            if result == 'win':
                weighted_score += 3 * weight
            elif result == 'draw':
                weighted_score += 1 * weight
            # Loss adds 0
        
        # Normalize to 0.5-1.5 range
        if total_weight > 0:
            form_score = weighted_score / (total_weight * 3)  # Divide by max possible score
            return 0.5 + form_score  # Scale to 0.5-1.5
        
        return 1.0

    def _get_result(self, home_goals, away_goals):
        if home_goals > away_goals:
            return "home_win"
        elif away_goals > home_goals:
            return "away_win"
        return "draw"

    def calculate_historical_factors(self, team1, team2):
        """Calculate factors based on historical matches"""
        key1 = f"{team1}_{team2}"
        key2 = f"{team2}_{team1}"
        
        total_matches = 0
        team1_wins = 0
        team2_wins = 0
        draws = 0
        team1_goals = 0
        team2_goals = 0
        
        for key in [key1, key2]:
            if key in self.historical_data:
                for match in self.historical_data[key]:
                    total_matches += 1
                    if match['result'] == 'home_win':
                        if key == key1:
                            team1_wins += 1
                        else:
                            team2_wins += 1
                    elif match['result'] == 'away_win':
                        if key == key1:
                            team2_wins += 1
                        else:
                            team1_wins += 1
                    else:
                        draws += 1
                    
                    if key == key1:
                        team1_goals += match['home_goals']
                        team2_goals += match['away_goals']
                    else:
                        team1_goals += match['away_goals']
                        team2_goals += match['home_goals']
        
        if total_matches == 0:
            return {
                'win_ratio': 0.5,
                'draw_ratio': 0.2,
                'goal_ratio': 1.0,
                'total_matches': 0
            }
        
        return {
            'win_ratio': team1_wins / total_matches,
            'draw_ratio': draws / total_matches,
            'goal_ratio': team1_goals / (team1_goals + team2_goals) if (team1_goals + team2_goals) > 0 else 1.0,
            'total_matches': total_matches
        }

    def calculate_team_strength(self, team_name, is_home=False, week=1):
        """Enhanced with current season momentum and form"""
        team = self.teams_data[team_name]
        
        # Calculate current season form (heavily weighted on first 4 matches)
        current_form = self.calculate_current_season_form(team_name, week)
        
        # Base strength (weighted average of attributes)
        base_strength = (
            team["attack"] * 0.4 +
            team["defense"] * 0.35 +
            team["midfield"] * 0.25
        )
        
        # Apply current season form and momentum
        strength = base_strength * current_form * team['season_momentum'] * team['confidence_boost']

        # Home advantage
        if is_home:
            strength *= team["home_advantage"]
        
        # Fatigue factor (increases over weeks)
        fatigue_penalty = 1 - (team["fatigue"] * week * 0.02)
        strength *= max(0.7, fatigue_penalty)
        
        # Injury impact
        injury_impact = 1 - (team["injury_rate"] * random.uniform(0, 1))
        strength *= injury_impact
        
        return max(30, strength)  # Minimum strength threshold

    def poisson_goal_probability(self, lambda_val, goals):
        """Calculate probability of scoring exactly 'goals' using Poisson distribution"""
        return (math.exp(-lambda_val) * (lambda_val ** goals)) / math.factorial(goals)

    def predict_match_advanced(self, home_team, away_team, week=1, weather="sunny", importance="regular", use_current_season=True):
        """Enhanced with current season momentum and early match impact"""
        
        # Automatically set importance for first 4 matches
        if week <= 4:
            importance = "early_season"
        
        # Get historical factors
        historical = self.calculate_historical_factors(home_team, away_team)

        # Get weather condition
        weather_factor = self.weather_impact.get(weather, 1.0)
        importance_factor = self.match_importance.get(importance, 1.0)
        
        # Calculate team strengths with current season factors
        home_strength = self.calculate_team_strength(home_team, is_home=True, week=week)
        away_strength = self.calculate_team_strength(away_team, is_home=False, week=week)
        
        # Apply external factors
        home_strength *= weather_factor * importance_factor
        away_strength *= weather_factor * importance_factor
        
        # Expected goals using Poisson model
        home_attack = self.teams_data[home_team]["attack"] / 100
        away_defense = self.teams_data[away_team]["defense"] / 100
        home_lambda = home_strength * home_attack * (1 - away_defense) * 0.03
        
        away_attack = self.teams_data[away_team]["attack"] / 100
        home_defense = self.teams_data[home_team]["defense"] / 100
        away_lambda = away_strength * away_attack * (1 - home_defense) * 0.025
        
        # Generate multiple scenarios and pick the most probable
        scenarios = []
        for home_goals in range(6):
            for away_goals in range(6):
                prob = (self.poisson_goal_probability(home_lambda, home_goals) * 
                       self.poisson_goal_probability(away_lambda, away_goals))
                scenarios.append((home_goals, away_goals, prob))
        
        # Sort by probability and add some randomness
        scenarios.sort(key=lambda x: x[2], reverse=True)
        
        # Pick from top 3 most probable scenarios with weighted randomness
        weights = [0.5, 0.3, 0.2]
        chosen_scenario = random.choices(scenarios[:3], weights=weights)[0]
        
        home_goals, away_goals = chosen_scenario[0], chosen_scenario[1]
        
        # Calculate win probability
        home_win_prob = sum(prob for h, a, prob in scenarios if h > a)
        draw_prob = sum(prob for h, a, prob in scenarios if h == a)
        away_win_prob = sum(prob for h, a, prob in scenarios if h < a)
        
        # Adjust probabilities based on historical data and current season momentum
        if historical['total_matches'] > 0:
            home_win_prob *= (1 + (historical['win_ratio'] - 0.5) * 0.2)
            away_win_prob *= (1 + (0.5 - historical['win_ratio']) * 0.2)
            draw_prob *= (1 + (historical['draw_ratio'] - 0.2) * 0.3)
        
        # Apply current season momentum to probabilities (stronger effect for matches after week 4)
        if week > 4:
            momentum_factor = 0.3  # Strong momentum effect after early matches
            home_momentum = self.teams_data[home_team]['season_momentum']
            away_momentum = self.teams_data[away_team]['season_momentum']
            
            home_win_prob *= (1 + (home_momentum - 1) * momentum_factor)
            away_win_prob *= (1 + (away_momentum - 1) * momentum_factor)
            
            # Normalize probabilities
            total = home_win_prob + away_win_prob + draw_prob
            home_win_prob /= total
            away_win_prob /= total
            draw_prob /= total

        # Determine result
        if home_goals > away_goals:
            result = f"{home_team} wins"
            confidence = home_win_prob
        elif away_goals > home_goals:
            result = f"{away_team} wins"
            confidence = away_win_prob
        else:
            result = "Draw"
            confidence = draw_prob
        
        return {
            "team1": home_team,
            "team2": away_team,
            "score1": home_goals,
            "score2": away_goals,
            "result": result,
            "confidence": round(confidence * 100, 1),
            "weather": weather,
            "match_importance": importance,
            "week": week,
            "home_strength": round(home_strength, 1),
            "away_strength": round(away_strength, 1),
            "current_season_factors": {
                "home_momentum": round(self.teams_data[home_team]['season_momentum'], 2),
                "away_momentum": round(self.teams_data[away_team]['season_momentum'], 2),
                "home_confidence": round(self.teams_data[home_team]['confidence_boost'], 2),
                "away_confidence": round(self.teams_data[away_team]['confidence_boost'], 2)
            },
            "expected_goals": {
                "home": round(home_lambda, 2),
                "away": round(away_lambda, 2)
            },
            "win_probabilities": {
                "home_win": round(home_win_prob * 100, 1),
                "draw": round(draw_prob * 100, 1),
                "away_win": round(away_win_prob * 100, 1)
            }
        }

    def simulate_season(self, num_simulations=1000):
        """Monte Carlo simulation with early match impact"""
        season_results = {}
        
        matchups = [
            ("Lions", "Tigers"), ("Bears", "Wolves"),
            ("Lions", "Bears"), ("Tigers", "Wolves"),
            ("Lions", "Wolves"), ("Tigers", "Bears")
        ]
        
        for _ in range(num_simulations):
            # Reset team states for each simulation
            temp_teams_data = {}
            for team, data in self.teams_data.items():
                temp_teams_data[team] = data.copy()
                temp_teams_data[team]['season_momentum'] = 1.0
                temp_teams_data[team]['confidence_boost'] = 1.0
            
            # Temporarily store original data
            original_data = self.teams_data
            self.teams_data = temp_teams_data
            
            season_points = {team: 0 for team in self.teams_data.keys()}
            temp_results = {}
            
            for week, (team1, team2) in enumerate(matchups, 1):
                weather = random.choice(self.weather_conditions)
                
                prediction = self.predict_match_advanced(team1, team2, week, weather)
                
                # Update points
                if prediction["score1"] > prediction["score2"]:
                    season_points[team1] += 3
                elif prediction["score2"] > prediction["score1"]:
                    season_points[team2] += 3
                else:
                    season_points[team1] += 1
                    season_points[team2] += 1
                
                # Update momentum for future matches in this simulation
                self._update_team_momentum(team1, team2, prediction["score1"], prediction["score2"], week)
            
            # Restore original data
            self.teams_data = original_data
            
            # Store results
            winner = max(season_points, key=season_points.get)
            if winner not in season_results:
                season_results[winner] = 0
            season_results[winner] += 1
        
        # Convert to probabilities
        for team in season_results:
            season_results[team] = round((season_results[team] / num_simulations) * 100, 1)
        
        return season_results

    def generate_tactical_analysis(self, team1, team2, week=1):
        """Generate tactical analysis with current season context"""
        team1_data = self.teams_data[team1]
        team2_data = self.teams_data[team2]
        
        analysis = {
            "key_battles": [],
            "tactical_advantages": {},
            "recommended_strategy": {},
            "momentum_analysis": {}
        }
        
        # Key battles
        if team1_data["attack"] > team2_data["defense"]:
            analysis["key_battles"].append(f"{team1}'s attack vs {team2}'s defense - Advantage: {team1}")
        else:
            analysis["key_battles"].append(f"{team1}'s attack vs {team2}'s defense - Advantage: {team2}")
        
        if team2_data["attack"] > team1_data["defense"]:
            analysis["key_battles"].append(f"{team2}'s attack vs {team1}'s defense - Advantage: {team2}")
        else:
            analysis["key_battles"].append(f"{team2}'s attack vs {team1}'s defense - Advantage: {team1}")
        
        # Current season momentum analysis
        if week > 4:
            analysis["momentum_analysis"] = {
                f"{team1}_momentum": "High" if team1_data['season_momentum'] > 1.2 else "Low" if team1_data['season_momentum'] < 0.8 else "Normal",
                f"{team2}_momentum": "High" if team2_data['season_momentum'] > 1.2 else "Low" if team2_data['season_momentum'] < 0.8 else "Normal",
                "early_season_impact": "High - First 4 matches heavily influencing current form"
            }
        else:
            analysis["momentum_analysis"] = {
                "match_importance": "Critical early season match - will heavily influence future performance"
            }
        
        # Tactical advantages
        if team1_data["midfield"] > team2_data["midfield"]:
            analysis["tactical_advantages"][team1] = "Midfield dominance"
        else:
            analysis["tactical_advantages"][team2] = "Midfield control"
        
        # Recommended strategies based on current form
        current_form1 = self.calculate_current_season_form(team1, week)
        current_form2 = self.calculate_current_season_form(team2, week)
        
        if current_form1 > 1.2:
            analysis["recommended_strategy"][team1] = "Maintain aggressive approach - riding high confidence"
        elif current_form1 < 0.8:
            analysis["recommended_strategy"][team1] = "Focus on defensive stability - rebuild confidence"
        else:
            analysis["recommended_strategy"][team1] = "Balanced approach"
        
        if current_form2 > 1.2:
            analysis["recommended_strategy"][team2] = "Maintain aggressive approach - riding high confidence"
        elif current_form2 < 0.8:
            analysis["recommended_strategy"][team2] = "Focus on defensive stability - rebuild confidence"
        else:
            analysis["recommended_strategy"][team2] = "Balanced approach"
        
        return analysis

def predict_league_outcomes():
    """Main prediction function with enhanced early match impact"""
    predictor = AdvancedPredictor()
    
    try:
        # Load historical data if available
        historical_matches = json.loads(sys.stdin.read())
        predictor.load_historical_data(historical_matches)
    except Exception as e:
        print(f"Warning: Could not load historical data: {str(e)}", file=sys.stderr)
    
    # Generate detailed predictions for upcoming matches
    upcoming_matches = [
        ("Lions", "Tigers"),
        ("Bears", "Wolves"),
        ("Lions", "Bears"),
        ("Tigers", "Wolves"),
        ("Lions", "Wolves"),
        ("Tigers", "Bears")
    ]
    
    predictions = []
    for week, (team1, team2) in enumerate(upcoming_matches, 1):
        weather = random.choice(predictor.weather_conditions)
        
        # Get detailed prediction
        prediction = predictor.predict_match_advanced(team1, team2, week, weather)
        
        # Add tactical analysis
        tactical_analysis = predictor.generate_tactical_analysis(team1, team2, week)
        prediction["tactical_analysis"] = tactical_analysis
        
        predictions.append(prediction)
        
        # Simulate this match result and update team momentum for future predictions
        # This creates the cascading effect where early results heavily influence later matches
        predictor.update_current_season_results({
            'home_team': team1,
            'away_team': team2,
            'home_goals': prediction["score1"],
            'away_goals': prediction["score2"],
            'week': week
        })
    
    # Add season simulation results
    season_simulation = predictor.simulate_season(1000)
    
    # Return comprehensive analysis
    return {
        "match_predictions": predictions,
        "season_simulation": {
            "championship_probabilities": season_simulation,
            "simulation_runs": 1000,
            "methodology": "Monte Carlo simulation with enhanced early match impact"
        },
        "prediction_metadata": {
            "algorithm": "Advanced Multi-Factor Prediction Model with Early Season Momentum",
            "factors_considered": [
                "Team attributes (attack, defense, midfield)",
                "Historical form and momentum",
                "Current season momentum (3x weight for first 4 matches)",
                "Confidence boost from early results",
                "Home advantage",
                "Weather conditions",
                "Match importance (enhanced for early season)",
                "Fatigue and injury factors",
                "Poisson distribution for goal probability"
            ],
            "early_match_impact": "First 4 matches have 3x impact on team momentum and form calculations",
            "confidence_level": "High",
            "last_updated": datetime.now().isoformat()
        }
    }

if __name__ == "__main__":
    try:
        result = predict_league_outcomes()
        print(json.dumps(result, indent=2))
    except Exception as e:
        error_response = {"error": str(e), "timestamp": datetime.now().isoformat()}
        print(json.dumps(error_response))