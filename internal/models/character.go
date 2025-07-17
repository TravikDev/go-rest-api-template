package models

// Character represents a player's avatar bound to a user.
type Character struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Nickname   string  `json:"nickname"`
	Level      int     `json:"level"`
	Experience int     `json:"experience"`
	PosX       float64 `json:"x"`
	PosY       float64 `json:"y"`
	PosZ       float64 `json:"z"`
}
