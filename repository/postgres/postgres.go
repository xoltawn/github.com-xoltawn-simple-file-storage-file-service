package postgres

import (
	"context"

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
	return
}
