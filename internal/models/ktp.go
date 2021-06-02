package models

import "time"

type Ktp struct {
	ID            int64
	Nama          string
	NIK           string
	Jenis_kelamin string
	Tanggal_lahir time.Time
	Alamat        string
	Agama         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Tablename is name of the table
func (Ktp) TableName() string {
	return "ktp"
}
