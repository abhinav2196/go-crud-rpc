package Item

import "github.com/jinzhu/gorm"

type Item struct{
	gorm.Model
	Title string `json:"title"`
	Body  string    `json:"body"`
}