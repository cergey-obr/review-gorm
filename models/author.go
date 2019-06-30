package models

type Author struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
