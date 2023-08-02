package model

import (
// "gorm.io/gorm"
)

type Matakuliah struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	NamaMK          string `gorm:"not null" json:"nama_mk"`
	SKS             uint   `gorm:"not null;check:sks > 0" json:"sks"`
	JurusanMK       string `gorm:"not null" json:"jurusan_mk"`
	SemesterMinimal uint   `gorm:"not null;check:semester_minimal > 0" json:"semester_minimal"`
	Prediksi        string `gorm:"not null;type:enum('A', 'AB', 'B', 'BC', 'C', 'D', 'E')" json:"prediksi"`
}
