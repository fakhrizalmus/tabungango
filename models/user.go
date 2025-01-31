package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama       string `json:"nama" db:"nama"`
	NIK        string `gorm:"unique" json:"nik" db:"nik"`
	NoHP       string `gorm:"unique" json:"no_hp" db:"no_hp"`
	NoRekening string `json:"no_rekening" db:"no_rekening"`
	Saldo      int64  `gorm:"default:0" json:"saldo" db:"saldo"`
}
