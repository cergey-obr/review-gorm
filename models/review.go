package models

type Review struct {
	Id        uint    `gorm:"primary_key" json:"id"`
	Title     string  `json:"title"`
	CreatedAt string  `json:"created_at"`
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
	Photos    []Photo `gorm:"foreignKey:ReviewId" json:"photos"`
}

type ProductReviews struct {
	Items  []Review `json:"items"`
	Totals Totals   `json:"totals"`
}

type Totals struct {
	Value  string `json:"value"`
	Count  string `json:"count"`
	Intval int    `json:"intval"`
}
