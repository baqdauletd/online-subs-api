package models

import "time"

type Sub struct{
	ID				string			`json:"id"  gorm:"type:uuid;  primaryKey"`
	ServiceName		string			`json:"service_name"  gorm:"not null"`
	Price			int 			`json:"price"  gorm:"not null"`
	UserID			string			`json:"user_id"  gorm:"type:uuid;  not null"`
	StartDate 		time.Time		`json:"start_date"  gorm:"not null"`
	EndDate			time.Time		`json:"end_date"  gorm:"not null"`
}