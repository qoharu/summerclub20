package model

import "time"

// Daily ... daily_updates table to store covid19 daily update
type Daily struct {
	ID        uint      `json:"id" gorm:"primary_key"` //this will be the primary key field
	Date      time.Time `json:"date" gorm:"type:DATE"`
	NewCase   int64     `json:"new_case"`
	TotalCase int64     `json:"total_case"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
