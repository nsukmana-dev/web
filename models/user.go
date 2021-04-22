package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nip          string
	NamaLengkap  string
	TempatLahir  string
	TanggalLahir time.Time
	Status       string
	Atasan       string
}
