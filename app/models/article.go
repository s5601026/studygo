package models

type Article struct {
	Id    uint64 `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
