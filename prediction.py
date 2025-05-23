import random
import json

# Takım güçleri: 1-10 arasında
team_strengths = {
    "TeamA": 8,
    "TeamB": 6,
    "TeamC": 5,
    "TeamD": 7
}

def predict_match(team1, team2):
    strength1 = team_strengths.get(team1, 5)
    strength2 = team_strengths.get(team2, 5)

    # Basit skor tahmini: güçlere göre rastgele
    score1 = random.randint(0, strength1)
    score2 = random.randint(0, strength2)

    if score1 > score2:
        result = f"{team1} wins"
    elif score2 > score1:
        result = f"{team2} wins"
    else:
        result = "Draw"

    return {"team1": team1, "team2": team2, "score1": score1, "score2": score2, "result": result}

if __name__ == "__main__":
    # Örnek kullanım
    matches = [
        ("TeamA", "TeamB"),
        ("TeamC", "TeamD")
    ]

    predictions = [predict_match(t1, t2) for t1, t2 in matches]
    print(json.dumps(predictions))
