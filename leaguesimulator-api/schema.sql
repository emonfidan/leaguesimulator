DROP DATABASE IF EXISTS leaguesimulator;
CREATE DATABASE leaguesimulator;
USE leaguesimulator;

CREATE TABLE teams (
    name VARCHAR(100) PRIMARY KEY,
    points INT DEFAULT 0,
    played INT DEFAULT 0,
    wins INT DEFAULT 0,
    draws INT DEFAULT 0,
    losses INT DEFAULT 0,
    goals_for INT DEFAULT 0,
    goals_against INT DEFAULT 0,
    strength INT NOT NULL
);


CREATE TABLE matches (
    id INT PRIMARY KEY AUTO_INCREMENT,
    week INT NOT NULL,
    home_team_name VARCHAR(100) NOT NULL,
    away_team_name VARCHAR(100) NOT NULL,
    home_goals INT DEFAULT 0,
    away_goals INT DEFAULT 0,
    played BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (home_team_name) REFERENCES teams(name),
    FOREIGN KEY (away_team_name) REFERENCES teams(name)
);

CREATE TABLE predictions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    team_name VARCHAR(100) NOT NULL,
    predicted_rank INT NOT NULL,
    week_submitted INT NOT NULL,
    FOREIGN KEY (team_name) REFERENCES teams(name)
);

CREATE TABLE historical_matches (
    id INT PRIMARY KEY AUTO_INCREMENT,
    season INT NOT NULL,
    week INT NOT NULL,
    home_team_name VARCHAR(100) NOT NULL,
    away_team_name VARCHAR(100) NOT NULL,
    home_goals INT DEFAULT 0,
    away_goals INT DEFAULT 0,
    FOREIGN KEY (home_team_name) REFERENCES teams(name),
    FOREIGN KEY (away_team_name) REFERENCES teams(name)
);
