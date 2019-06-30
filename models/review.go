package models

type Review struct {
	Id        uint    `gorm:"primary_key" json:"id"`
	Title     string  `json:"title"`
	ProductId int     `json:"productId"`
	StatusId  int     `json:"statusId"`
	AuthorId  int     `json:"-"`
	Detail    string  `json:"detail"`
	Recommend bool    `json:"recommend"`
	Website   int     `json:"website"`
	Overall   int     `json:"overall"`
	Quality   int     `json:"quality"`
	Fit       int     `json:"fit"`
	Style     int     `json:"style"`
	Author    Author  `gorm:"foreignKey:AuthorId" json:"author"`
	Photos    []Photo `gorm:"foreignkey:ReviewId" json:"photos"`
}
