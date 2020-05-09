package model

import "time"

// Daily ... dailies table to store covid19 daily update
type Daily struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"` //this will be the primary key field
	Date      time.Time `json:"date" binding:"required" gorm:"type:DATE;NOT NULL;UNIQUE"`
	NewCase   int64     `json:"new_case" binding:"required" gorm:"NOT NULL"`
	TotalCase int64     `json:"total_case" binding:"required" gorm:"NOT NULL"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
