package models

import "time"

type Room struct {
    ID           string    `json:"id"`
    TeamHome     string    `json:"team_home"`
    TeamAway     string    `json:"team_away"`
    Status       string    `json:"status"`
    MatchDate    time.Time `json:"match_date"`
    CreatedAt    time.Time `json:"created_at"`
    ScoreHome    *int      `json:"score_home"`
    ScoreAway    *int      `json:"score_away"`
}

type Competition struct {
    ID                string    `json:"id"`
    Name              string    `json:"name"`
    Type              string    `json:"type"`
    Country           string    `json:"country"`
    Status            string    `json:"status"`
    Teams             []string  `json:"teams"`
    MatchDate         time.Time `json:"match_date"`
    CreatedAt         time.Time `json:"created_at"`
    TotalRooms        int       `json:"total_rooms"`
    TotalParticipants int       `json:"total_participants"`
}

type Match struct {
    ID        string    `json:"id"`
    TeamHome  string    `json:"team_home"`
    TeamAway  string    `json:"team_away"`
    Status    string    `json:"status"`
    MatchDate time.Time `json:"match_date"`
    ScoreHome *int      `json:"score_home"`
    ScoreAway *int      `json:"score_away"`
    CreatedAt time.Time `json:"created_at"`
}