package handlers

import (
	"context"
	"mariadb/configs/database"
	"mariadb/graph/model"
	"mariadb/internal/models"
	"time"
)

func stringToDate(value string) time.Time {
	var layoutFormat = "2006-01-02 15:04:05"
	var date, _ = time.Parse(layoutFormat, value)

	return date
}

func CreateKtpHandler(ctx context.Context, input model.KtpBody) (*models.Ktp, error) {
	ktp := models.Ktp{
		NIK:           input.Nik,
		Nama:          input.Nama,
		Jenis_kelamin: input.JenisKelamin,
		Tanggal_lahir: stringToDate(input.TanggalLahir),
		Alamat:        input.Alamat,
		Agama:         input.Agama,
		CreatedAt:     time.Now(),
	}

	res := database.DB.Create(&ktp) // pass pointer of data to Create

	if res.Error != nil {
		return nil, res.Error
	}

	return &ktp, nil
}

func DeleteKtpHandler(ctx context.Context, id int64) (bool, error) {
	res := database.DB.Delete(&models.Ktp{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

func EditKtpHandler(ctx context.Context, id int64, input model.KtpBody) (*models.Ktp, error) {
	// Assumption where update all fields get updated
	newKtp := models.Ktp{
		NIK:           input.Nik,
		Nama:          input.Nama,
		Jenis_kelamin: input.JenisKelamin,
		Tanggal_lahir: stringToDate(input.TanggalLahir),
		Alamat:        input.Alamat,
		Agama:         input.Agama,
		UpdatedAt:     time.Now(),
	}
	res := database.DB.Model(&models.Ktp{}).Where("id = ?", id).Updates(newKtp)

	if res.Error != nil {
		return nil, res.Error
	}

	return &newKtp, nil
}

func GetAllKtpHandler(ctx context.Context) ([]*models.Ktp, error) {
	var ktp []*models.Ktp
	res := database.DB.Find(&ktp)

	if res.Error != nil {
		return nil, res.Error
	}

	return ktp, nil
}
