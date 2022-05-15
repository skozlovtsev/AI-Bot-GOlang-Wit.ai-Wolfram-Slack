package models

type User struct {
	Mail     string `gorm:"primaryKey" json:"mail"`
	Passhash string `gorm:"not null" json:"passhash"`
}