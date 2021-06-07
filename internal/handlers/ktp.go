package handlers

import (
	"context"
	"mariadb/configs/database"
	"mariadb/graph/model"
	"mariadb/internal/helpers"
	"mariadb/internal/models"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// Function for create new ktp
// This function will return ktp object
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

// Function for delete ktp by id
// This function will return boolean
func DeleteKtpHandler(ctx context.Context, id int64) (bool, error) {
	res := database.DB.Delete(&models.Ktp{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

// Function for edit ktp by id
// This function will return new edit ktp
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

// This function for get all ktp
// This function will return all ktp
func GetAllKtpHandler(ctx context.Context) ([]*models.Ktp, error) {
	var ktp []*models.Ktp
	res := database.DB.Find(&ktp)

	if res.Error != nil {
		return nil, res.Error
	}

	return ktp, nil
}

func GetPaginationKtpHandler(ctx context.Context, input model.Pagination) (*model.PaginationResultKtp, error) {
	var ktp []*models.Ktp
	var ktpEdges []*model.PaginationEdgeKtp
	var total, startId int64

	if input.After != nil {
		startId = *input.After
	} else {
		startId = 1
	}

	// Query for generate pagination
	paginateQuery, paginateArgs, paginateSqlErr :=
		sq.Select("*").From("ktp").Where("id >= ?", startId+input.Offset).OrderBy("id ASC").Limit(uint64(input.First)).ToSql()
	if paginateSqlErr != nil {
		return nil, paginateSqlErr
	}

	// Query for find total
	idTotalQuery, _, idTotalSqlErr := sq.Select("COUNT(id)").From("ktp").ToSql()
	if idTotalSqlErr != nil {
		return nil, idTotalSqlErr

	}

	database.DB.Raw(paginateQuery, paginateArgs...).Scan(&ktp)
	database.DB.Raw(idTotalQuery).Scan(&total)

	// Generate ktp edges
	for _, val := range ktp {
		ktpEdges = append(ktpEdges, &model.PaginationEdgeKtp{
			Node:   val,
			Cursor: val.ID,
		})
	}

	// Generate pagination result model
	ktpPaginationResult := model.PaginationResultKtp{
		TotalCount: total,
		Edges:      ktpEdges,
		PageInfo: &model.PaginationInfo{
			EndCursor:   ktpEdges[int64(len(ktpEdges))-1].Cursor,
			HasNextPage: ktpEdges[int64(len(ktpEdges))-1].Node.ID < total,
		},
	}

	return &ktpPaginationResult, nil
}
