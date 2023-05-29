package repository

import (
	"context"
	"fp-mbd-amidrive/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MobilRepository interface {
	AddMobil(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error)
	GetAllMobil(ctx context.Context) ([]entity.Mobil, error)
	FindMobilByID(ctx context.Context, mobilID uuid.UUID) (entity.Mobil, error)
	DeleteMobil(ctx context.Context, mobilID uuid.UUID) error
	UpdateMobil(ctx context.Context, mobil entity.Mobil) error
}

type mobilConnection struct {
	connection *gorm.DB
}

func NewMobilRepository(db *gorm.DB) MobilRepository {
	return &mobilConnection{
		connection: db,
	}
}

func (db *mobilConnection) AddMobil(ctx context.Context, mobil entity.Mobil) (entity.Mobil, error) {
	mobil.ID = uuid.New()
	uc := db.connection.Create(&mobil)
	if uc.Error != nil {
		return entity.Mobil{}, uc.Error
	}
	return mobil, nil
}

func (db *mobilConnection) GetAllMobil(ctx context.Context) ([]entity.Mobil, error) {
	var listMobil []entity.Mobil
	tx := db.connection.Find(&listMobil)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listMobil, nil
}

func (db *mobilConnection) GetMobil(ctx context.Context, MobilID uuid.UUID) ([]entity.Mobil, error) {
	var listMobil []entity.Mobil
	tx := db.connection.Find(&listMobil)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listMobil, nil
}

func (db *mobilConnection) FindMobilByEmail(ctx context.Context, email string) (entity.Mobil, error) {
	var mobil entity.Mobil
	ux := db.connection.Where("email = ?", email).Take(&mobil)
	if ux.Error != nil {
		return mobil, ux.Error
	}
	return mobil, nil
}

func (db *mobilConnection) FindMobilByID(ctx context.Context, mobilID uuid.UUID) (entity.Mobil, error) {
	var mobil entity.Mobil
	ux := db.connection.Where("id = ?", mobilID).Take(&mobil)
	if ux.Error != nil {
		return mobil, ux.Error
	}
	return mobil, nil
}

func (db *mobilConnection) DeleteMobil(ctx context.Context, mobilID uuid.UUID) error {
	uc := db.connection.Delete(&entity.Mobil{}, &mobilID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *mobilConnection) UpdateMobil(ctx context.Context, mobil entity.Mobil) error {
	uc := db.connection.Updates(&mobil)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
