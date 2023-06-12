package repository

import (
	"context"
	"fp-mbd-amidrive/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MetodePembayaranRepository interface {
	AddMetodePembayaran(ctx context.Context, metodePembayaran entity.MetodePembayaran) (entity.MetodePembayaran, error)
	GetAllMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error)
	FindMetodePembayaranByID(ctx context.Context, metodePembayaranID uuid.UUID) (entity.MetodePembayaran, error)
	DeleteMetodePembayaran(ctx context.Context, metodePembayaranID uuid.UUID) error
	UpdateMetodePembayaran(ctx context.Context, metodePembayaran entity.MetodePembayaran) error
}

type metodePembayaranConnection struct {
	connection *gorm.DB
}

func NewMetodePembayaranRepository(db *gorm.DB) MetodePembayaranRepository {
	return &metodePembayaranConnection{
		connection: db,
	}
}

func (db *metodePembayaranConnection) AddMetodePembayaran(ctx context.Context, metodePembayaran entity.MetodePembayaran) (entity.MetodePembayaran, error) {
	uc := db.connection.Create(&metodePembayaran)
	if uc.Error != nil {
		return entity.MetodePembayaran{}, uc.Error
	}
	return metodePembayaran, nil
}

func (db *metodePembayaranConnection) GetAllMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error) {
	var listMetodePembayaran []entity.MetodePembayaran
	tx := db.connection.Find(&listMetodePembayaran)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listMetodePembayaran, nil
}

func (db *metodePembayaranConnection) GetMetodePembayaran(ctx context.Context, MetodePembayaranID uuid.UUID) ([]entity.MetodePembayaran, error) {
	var listMetodePembayaran []entity.MetodePembayaran
	tx := db.connection.Find(&listMetodePembayaran)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listMetodePembayaran, nil
}

func (db *metodePembayaranConnection) FindMetodePembayaranByEmail(ctx context.Context, email string) (entity.MetodePembayaran, error) {
	var metodePembayaran entity.MetodePembayaran
	ux := db.connection.Where("email = ?", email).Take(&metodePembayaran)
	if ux.Error != nil {
		return metodePembayaran, ux.Error
	}
	return metodePembayaran, nil
}

func (db *metodePembayaranConnection) FindMetodePembayaranByID(ctx context.Context, metodePembayaranID uuid.UUID) (entity.MetodePembayaran, error) {
	var metodePembayaran entity.MetodePembayaran
	ux := db.connection.Where("id = ?", metodePembayaranID).Take(&metodePembayaran)
	if ux.Error != nil {
		return metodePembayaran, ux.Error
	}
	return metodePembayaran, nil
}

func (db *metodePembayaranConnection) DeleteMetodePembayaran(ctx context.Context, metodePembayaranID uuid.UUID) error {
	uc := db.connection.Delete(&entity.MetodePembayaran{}, &metodePembayaranID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *metodePembayaranConnection) UpdateMetodePembayaran(ctx context.Context, metodePembayaran entity.MetodePembayaran) error {
	uc := db.connection.Updates(&metodePembayaran)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
