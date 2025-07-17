package models

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	PasswordHash   string `json:"-"`
	FailedAttempts int    `json:"-"`
	Locked         bool   `json:"locked"`
}
