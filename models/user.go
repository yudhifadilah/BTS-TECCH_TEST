package models

import (
	"database/sql"
)

type User struct {
	ID        uint         `gorm:"primaryKey"`
	Nama      string       `json:"nama" gorm:"column:nama;not null"`
	Username  string       `json:"username" gorm:"unique;not null"`
	Password  string       `json:"password" gorm:"not null"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
