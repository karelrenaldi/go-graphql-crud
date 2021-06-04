package handlers

import (
	"context"
	"mariadb/configs/database"
	"mariadb/graph/model"
	"mariadb/internal/helpers"
	"mariadb/internal/models"
	"time"
)

/*
	@desc Function for create new ktp data

	@param ctx(context.Context) : context
	@param input(model.KtpBody) : input value for new ktp data

	@output (*models.Ktp) : data value
	@output (error) : error
*/
func CreateKtpHandler(ctx context.Context, input model.KtpBody) (*models.Ktp, error) {
	ktp := models.Ktp{
		NIK:           input.Nik,
		Nama:          input.Nama,
		Jenis_kelamin: input.JenisKelamin,
		Tanggal_lahir: helpers.StringToDate(input.TanggalLahir),
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

/*
	@desc Function for delete data by id

	@param id(int64) : id's data
	@param ctx(context.Context) : context

	@output (bool) : boolean value for delete status
	@output (error) : error
*/
func DeleteKtpHandler(ctx context.Context, id int64) (bool, error) {
	res := database.DB.Delete(&models.Ktp{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

/*
	@desc Function for create new ktp data

	@param ctx(context.Context) : context
	@param id(int64) : id's data
	@param input(model.KtpBody) : input value for update ktp data

	@output (*models.Ktp) : updated data value
	@output (error) : error
*/
func EditKtpHandler(ctx context.Context, id int64, input model.KtpBody) (*models.Ktp, error) {
	// Assumption where update all fields get updated
	newKtp := models.Ktp{
		NIK:           input.Nik,
		Nama:          input.Nama,
		Jenis_kelamin: input.JenisKelamin,
		Tanggal_lahir: helpers.StringToDate(input.TanggalLahir),
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

/*
	@desc Function for get all ktp data

	@param ctx(context.Context) : context

	@output ([]*models.Ktp) : list of data value
	@output (error) : error
*/
func GetAllKtpHandler(ctx context.Context) ([]*models.Ktp, error) {
	var ktp []*models.Ktp
	res := database.DB.Find(&ktp)

	if res.Error != nil {
		return nil, res.Error
	}

	return ktp, nil
}
