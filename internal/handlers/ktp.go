package handlers

import (
	"context"
	"mariadb/configs/database"
	"mariadb/graph/model"
	"mariadb/internal/helpers"
	"mariadb/internal/models"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// CreateKtpHandler is function for create new ktp
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

// DeleteKtpHandler is function for delete ktp by id
// This function will return boolean
func DeleteKtpHandler(ctx context.Context, id int64) (bool, error) {
	res := database.DB.Delete(&models.Ktp{}, id)

	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}

// EditKtpHandler is unction for edit ktp by id
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

// GetAllKtpHandler is function for get all ktp
// This function will return all ktp
func GetAllKtpHandler(ctx context.Context) ([]*models.Ktp, error) {
	var ktp []*models.Ktp
	res := database.DB.Find(&ktp)

	if res.Error != nil {
		return nil, res.Error
	}

	return ktp, nil
}

// GetPaginationKtpHandler is function for get ktp (pagination)
// This function will return ktp (pagination)
func GetPaginationKtpHandler(ctx context.Context, input model.Pagination) (*model.PaginationResultKtp, error) {
	var ktp []*models.Ktp
	var ktpEdges []*model.PaginationEdgeKtp
	var total, startId int64
	var nameArgs, nikArgs string

	// All row
	query := sq.Select("*").From("ktp")
	queryCount := sq.Select("COUNT(id)").From("ktp")

	// Check offset
	if input.After != nil {
		startId = *input.After + input.Offset
	} else {
		startId = 1 + input.Offset
	}

	// Check query input
	queryInput := strings.Split(input.Query, "&")
	for _, val := range queryInput {
		if strings.Contains(val, "nama") {
			nameArgs = strings.Split(val, "=")[1]
		} else if strings.Contains(val, "nik") {
			nikArgs = strings.Split(val, "=")[1]
		}
	}

	if len(nameArgs) != 0 && len(nikArgs) != 0 {
		query = query.Where("id >= ? AND (nama = ? OR nik = ?)", startId, nameArgs, nikArgs)
		queryCount = queryCount.Where("id >= ? AND (nama = ? OR nik = ?)", startId, nameArgs, nikArgs)
	} else if len(nameArgs) != 0 {
		query = query.Where("id >= ? AND nama = ?", startId, nameArgs)
		queryCount = queryCount.Where("id >= ? AND nama = ?", startId, nameArgs)
	} else if len(nikArgs) == 0 {
		query = query.Where("id >= ? AND nik = ?", startId, nikArgs)
		queryCount = queryCount.Where("id >= ? AND nik = ?", startId, nikArgs)
	} else {
		query = query.Where("id >= ?", startId)
		queryCount = queryCount.Where("id >= ?", startId)
	}

	// Check limit
	if input.First != 0 {
		query = query.Limit(uint64(input.First))
		queryCount = queryCount.Limit(uint64(input.First))
	} else {
		query = query.Limit(10)
		queryCount = queryCount.Limit(10)
	}

	// Check sort
	for _, val := range input.Sort {
		desc := strings.HasPrefix(val, "-")
		if desc {
			query = query.OrderBy(strings.Replace(val, "-", "", 1) + " desc")
		} else {
			query = query.OrderBy(val + " asc")
		}
	}

	// Generate sql query
	paginateQuery, paginateArgs, paginateSqlErr := query.ToSql()
	if paginateSqlErr != nil {
		return nil, paginateSqlErr
	}

	// Query for find total
	idTotalQuery, idTotalArgs, idTotalSqlErr := queryCount.ToSql()
	if idTotalSqlErr != nil {
		return nil, idTotalSqlErr
	}

	// Execute query command
	rs := database.DB.Raw(paginateQuery, paginateArgs...).Scan(&ktp)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// Execute query count command
	rs = database.DB.Raw(idTotalQuery, idTotalArgs...).Scan(&total)
	if rs.Error != nil {
		return nil, rs.Error
	}

	// Generate ktp edges
	for _, val := range ktp {
		ktpEdges = append(ktpEdges, &model.PaginationEdgeKtp{
			Node:   val,
			Cursor: helpers.EncodeCursor(val.ID),
		})
	}

	// End cursor handle
	var endCursor string
	if int64(len(ktpEdges)) == 0 {
		endCursor = ""
	} else {
		endCursor = ktpEdges[int64(len(ktpEdges))-1].Cursor
	}

	// Generate pagination result model
	ktpPaginationResult := model.PaginationResultKtp{
		TotalCount: total,
		Edges:      ktpEdges,
		PageInfo: &model.PaginationInfo{
			EndCursor:   endCursor,
			HasNextPage: int64(len(ktpEdges)) == int64(input.First),
		},
	}

	return &ktpPaginationResult, nil
}
