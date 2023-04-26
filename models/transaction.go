package models

import (
	"time"
)

type Transaction struct {
	ID    int       `gorm:"primary_key;auto_increment;not null"`
	Name  string    `gorm:"not null"`
	Type  int       `gorm:"not null;default:0;check: type >= 0 AND type <= 1"`
	Value int       `gorm:"not null;default:0"`
	Date  time.Time `gorm:"precision:3"`
}
