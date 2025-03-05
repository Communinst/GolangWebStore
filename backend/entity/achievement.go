package entities

type Achievement struct {
	AchievementId int    `json:"achievement_id" db:"achievement_id"`
	GameId        int    `json:"game_id" db:"game_id"`
	Name          string `json:"name" db:"name"`
	Description   string `json:"description" db:"description"`
}

type AchievementFrequency struct {
	AchievementId       int     `json:"achievement_id" db:"achievement_id"`
	FrequencyPercentage float32 `json:"frequency_percentage" db:"frequency_percentage"`
}
