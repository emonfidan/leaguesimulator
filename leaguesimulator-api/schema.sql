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
    strength INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE matches (
    id INT PRIMARY KEY AUTO_INCREMENT,
    week INT NOT NULL,
    home_team_name VARCHAR(100) NOT NULL,
    away_team_name VARCHAR(100) NOT NULL,
    home_goals INT DEFAULT 0,
    away_goals INT DEFAULT 0,
    played BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_week (week),
    INDEX idx_teams (home_team_name, away_team_name),
    CONSTRAINT fk_home_team FOREIGN KEY (home_team_name) REFERENCES teams(name) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_away_team FOREIGN KEY (away_team_name) REFERENCES teams(name) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE predictions (
    id INT PRIMARY KEY AUTO_INCREMENT,
    team_name VARCHAR(100) NOT NULL,
    predicted_rank INT NOT NULL,
    week_submitted INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (team_name) REFERENCES teams(name) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE historical_matches (
    id INT PRIMARY KEY AUTO_INCREMENT,
    season INT NOT NULL,
    week INT NOT NULL,
    home_team_name VARCHAR(100) NOT NULL,
    away_team_name VARCHAR(100) NOT NULL,
    home_goals INT DEFAULT 0,
    away_goals INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_season_week (season, week),
    FOREIGN KEY (home_team_name) REFERENCES teams(name) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (away_team_name) REFERENCES teams(name) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Insert default teams
INSERT INTO teams (name, strength) VALUES 
('Lions', 90),
('Tigers', 80),
('Bears', 70),
('Wolves', 60);