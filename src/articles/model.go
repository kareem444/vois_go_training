package articles

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Url           string `json:"url" gorm:"not null;"`
	Title         string `json:"title" gorm:"not null;varchar(255)"`
	TitleSelector string `json:"title_selector" gorm:"not null;varchar(255);"`
	UserID        string `json:"user_id" gorm:"not null;varchar(255)"`
}
