package models

type Photo struct {
	Id       int    `json:"id"`
	ReviewId int    `json:"-"`
	Url      string `json:"url"`
}
