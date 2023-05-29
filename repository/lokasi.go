package repository

import (
	"context"
	"fp-mbd-amidrive/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LokasiRepository interface {
	AddLokasi(ctx context.Context, lokasi entity.Lokasi) (entity.Lokasi, error)
	GetAllLokasi(ctx context.Context) ([]entity.Lokasi, error)
	FindLokasiByID(ctx context.Context, lokasiID uuid.UUID) (entity.Lokasi, error)
	DeleteLokasi(ctx context.Context, lokasiID uuid.UUID) error
	UpdateLokasi(ctx context.Context, lokasi entity.Lokasi) error
}

type lokasiConnection struct {
	connection *gorm.DB
}

func NewLokasiRepository(db *gorm.DB) LokasiRepository {
	return &lokasiConnection{
		connection: db,
	}
}

func (db *lokasiConnection) AddLokasi(ctx context.Context, lokasi entity.Lokasi) (entity.Lokasi, error) {
	lokasi.ID = uuid.New()
	uc := db.connection.Create(&lokasi)
	if uc.Error != nil {
		return entity.Lokasi{}, uc.Error
	}
	return lokasi, nil
}

func (db *lokasiConnection) GetAllLokasi(ctx context.Context) ([]entity.Lokasi, error) {
	var listLokasi []entity.Lokasi
	tx := db.connection.Find(&listLokasi)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listLokasi, nil
}

func (db *lokasiConnection) GetLokasi(ctx context.Context, LokasiID uuid.UUID) ([]entity.Lokasi, error) {
	var listLokasi []entity.Lokasi
	tx := db.connection.Find(&listLokasi)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listLokasi, nil
}

func (db *lokasiConnection) FindLokasiByEmail(ctx context.Context, email string) (entity.Lokasi, error) {
	var lokasi entity.Lokasi
	ux := db.connection.Where("email = ?", email).Take(&lokasi)
	if ux.Error != nil {
		return lokasi, ux.Error
	}
	return lokasi, nil
}

func (db *lokasiConnection) FindLokasiByID(ctx context.Context, lokasiID uuid.UUID) (entity.Lokasi, error) {
	var lokasi entity.Lokasi
	ux := db.connection.Where("id = ?", lokasiID).Take(&lokasi)
	if ux.Error != nil {
		return lokasi, ux.Error
	}
	return lokasi, nil
}

func (db *lokasiConnection) DeleteLokasi(ctx context.Context, lokasiID uuid.UUID) error {
	uc := db.connection.Delete(&entity.Lokasi{}, &lokasiID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *lokasiConnection) UpdateLokasi(ctx context.Context, lokasi entity.Lokasi) error {
	uc := db.connection.Updates(&lokasi)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
