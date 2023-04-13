package postgres

import (
	"context"
	"log"

	domain "github.com/xoltawn/simple-file-storage-file-service/domain"
	"gorm.io/gorm"
)

type filePostgresRepository struct {
	db *gorm.DB
}

func NewFilePostgresRepository(db *gorm.DB) *filePostgresRepository {
	return &filePostgresRepository{
		db: db,
	}
}

func (r *filePostgresRepository) SaveFile(ctx context.Context, fileInfo *domain.File) (err error) {
	result := r.db.Create(fileInfo)
	if result.RowsAffected == 0 {
		return domain.ErrInsertingRecord
	}

	if result.Error != nil {
		log.Println(result.Error.Error())
		return domain.ErrInsertingRecord
	}

	return
}

func (r *filePostgresRepository) SaveMutltipleFiles(ctx context.Context, files []*domain.File) (err error) {
	result := r.db.Create(&files)
	if result.RowsAffected == 0 {
		return domain.ErrInsertingRecord
	}

	if result.Error != nil {
		log.Println(result.Error.Error())
		return domain.ErrInsertingRecord
	}

	return
}

func (r *filePostgresRepository) FetchFiles(ctx context.Context, limit, offset int) (files []domain.File, err error) {
	tx := r.db.Limit(limit).Offset(offset).Find(&files)
	return files, tx.Error
}
