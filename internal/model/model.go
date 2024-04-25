package model

import "time"

type (
	Book struct {
		ID              uint      `json:"id" gorm:"primary_key;auto_increment"`
		Title           string    `json:"title" gorm:"type:varchar(255);not null"`
		Author          string    `json:"author" gorm:"type:varchar(255);not null"`
		PublicationDate time.Time `json:"publication_date" gorm:"type:date;not null"`
		ISBN            string    `json:"isbn" gorm:"type:varchar(20);not null;unique"`
	}

	User struct {
		ID       uint   `json:"id" gorm:"primary_key;auto_increment"`
		Email    string `json:"email" gorm:"type:varchar(255);not null;unique"`
		Password string `json:"password" gorm:"type:varchar(255);not null"`
		Name     string `json:"name" gorm:"type:varchar(255);not null"`
	}

	Progress struct {
		ID               uint `json:"id" gorm:"primary_key;auto_increment"`
		UserID           uint `json:"user_id" gorm:"not null"`
		BookID           uint `json:"book_id" gorm:"not null"`
		PageNumber       int  `json:"page_number" gorm:"not null"`
		CompletionStatus bool `json:"completion_status" gorm:"not null"`
	}
)
