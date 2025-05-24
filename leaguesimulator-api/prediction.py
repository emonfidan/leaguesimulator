import random
import json
import math
import sys
from datetime import datetime

class AdvancedPredictor:
    def __init__(self):
        # Enhanced team data with multiple attributes
        self.teams_data = {
            "Lions": {
                "attack": 90,
                "defense": 85,
                "midfield": 88,
                "form": 0.8,  # Recent form factor
                "home_advantage": 1.15,
                "fatigue": 0.0,
                "injury_rate": 0.05,
                "historical_performance": []
            },
            "Tigers": {
                "attack": 82,
                "defense": 78,
                "midfield": 80,
                "form": 0.75,
                "home_advantage": 1.12,
                "fatigue": 0.0,
                "injury_rate": 0.08,
                "historical_performance": []
            },
            "Bears": {
                "attack": 75,
                "defense": 80,
                "midfield": 72,
                "form": 0.7,
                "home_advantage": 1.1,
                "fatigue": 0.0,
                "injury_rate": 0.1,
                "historical_performance": []
            },
            "Wolves": {
                "attack": 68,
                "defense": 70,
                "midfield": 65,
                "form": 0.65,
                "home_advantage": 1.08,
                "fatigue": 0.0,
                "injury_rate": 0.12,
                "historical_performance": []
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
        
        # Match importance factor
        self.match_importance = {
            "regular": 1.0,
            "derby": 1.2,
            "final_week": 1.3
        }

    def calculate_team_strength(self, team_name, is_home=False, week=1):
        """Calculate dynamic team strength based on multiple factors"""
        team = self.teams_data[team_name]
        
        # Base strength (weighted average of attributes)
        base_strength = (
            team["attack"] * 0.4 +
            team["defense"] * 0.35 +
            team["midfield"] * 0.25
        )
        
        # Apply form factor
        strength = base_strength * team["form"]
        
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

    def predict_match_advanced(self, home_team, away_team, week=1, weather="sunny", importance="regular"):
        """Advanced match prediction using multiple algorithms"""
        
        # Get weather condition
        weather_factor = self.weather_impact.get(weather, 1.0)
        importance_factor = self.match_importance.get(importance, 1.0)
        
        # Calculate team strengths
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
            "home_strength": round(home_strength, 1),
            "away_strength": round(away_strength, 1),
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
        """Monte Carlo simulation of entire season"""
        season_results = {}
        
        matchups = [
            ("Lions", "Tigers"), ("Bears", "Wolves"),
            ("Lions", "Bears"), ("Tigers", "Wolves"),
            ("Lions", "Wolves"), ("Tigers", "Bears")
        ]
        
        for _ in range(num_simulations):
            season_points = {team: 0 for team in self.teams_data.keys()}
            
            for week, (team1, team2) in enumerate(matchups, 1):
                weather = random.choice(self.weather_conditions)
                importance = "final_week" if week >= 3 else "regular"
                
                prediction = self.predict_match_advanced(team1, team2, week, weather, importance)
                
                if prediction["score1"] > prediction["score2"]:
                    season_points[team1] += 3
                elif prediction["score2"] > prediction["score1"]:
                    season_points[team2] += 3
                else:
                    season_points[team1] += 1
                    season_points[team2] += 1
            
            # Store results
            winner = max(season_points, key=season_points.get)
            if winner not in season_results:
                season_results[winner] = 0
            season_results[winner] += 1
        
        # Convert to probabilities
        for team in season_results:
            season_results[team] = round((season_results[team] / num_simulations) * 100, 1)
        
        return season_results

    def generate_tactical_analysis(self, team1, team2):
        """Generate tactical analysis for the match"""
        team1_data = self.teams_data[team1]
        team2_data = self.teams_data[team2]
        
        analysis = {
            "key_battles": [],
            "tactical_advantages": {},
            "recommended_strategy": {}
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
        
        # Tactical advantages
        if team1_data["midfield"] > team2_data["midfield"]:
            analysis["tactical_advantages"][team1] = "Midfield dominance"
        else:
            analysis["tactical_advantages"][team2] = "Midfield control"
        
        # Recommended strategies
        if team1_data["attack"] > 80:
            analysis["recommended_strategy"][team1] = "Aggressive attacking play"
        elif team1_data["defense"] > 80:
            analysis["recommended_strategy"][team1] = "Defensive counter-attack"
        else:
            analysis["recommended_strategy"][team1] = "Balanced approach"
        
        return analysis

def predict_league_outcomes():
    """Main prediction function with enhanced features"""
    predictor = AdvancedPredictor()
    
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
        importance = "final_week" if week >= 3 else "regular"
        
        # Get detailed prediction
        prediction = predictor.predict_match_advanced(team1, team2, week, weather, importance)
        
        # Add tactical analysis
        tactical_analysis = predictor.generate_tactical_analysis(team1, team2)
        prediction["tactical_analysis"] = tactical_analysis
        
        predictions.append(prediction)
    
    # Add season simulation results
    season_simulation = predictor.simulate_season(1000)
    
    # Return comprehensive analysis
    return {
        "match_predictions": predictions,
        "season_simulation": {
            "championship_probabilities": season_simulation,
            "simulation_runs": 1000,
            "methodology": "Monte Carlo simulation with Poisson distribution"
        },
        "prediction_metadata": {
            "algorithm": "Advanced Multi-Factor Prediction Model",
            "factors_considered": [
                "Team attributes (attack, defense, midfield)",
                "Form and momentum",
                "Home advantage",
                "Weather conditions",
                "Match importance",
                "Fatigue and injury factors",
                "Poisson distribution for goal probability"
            ],
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