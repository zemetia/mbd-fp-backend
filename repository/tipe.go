package repository

import (
	"context"
	"fp-mbd-amidrive/entity"

	"gorm.io/gorm"
)

type TipeRepository interface {
	AddTipeMobil(ctx context.Context, tipeMobil entity.TipeMobil) (entity.TipeMobil, error)
	GetAllTipeMobil(ctx context.Context) ([]entity.TipeMobil, error)
	GetTipeMobilByID(ctx context.Context, tipeMobilID uint8) (entity.TipeMobil, error)
	DeleteTipeMobil(ctx context.Context, tipeMobil uint8) error
	UpdateTipeMobil(ctx context.Context, tipeMobil entity.TipeMobil) error

	AddTipeMesin(ctx context.Context, tipeMesin entity.TipeMesin) (entity.TipeMesin, error)
	GetAllTipeMesin(ctx context.Context) ([]entity.TipeMesin, error)
	GetTipeMesinByID(ctx context.Context, tipeMesinID uint8) (entity.TipeMesin, error)
	DeleteTipeMesin(ctx context.Context, tipeMesin uint8) error
	UpdateTipeMesin(ctx context.Context, tipeMesin entity.TipeMesin) error

	AddTipePersneling(ctx context.Context, tipeMobil entity.TipePersneling) (entity.TipePersneling, error)
	GetAllTipePersneling(ctx context.Context) ([]entity.TipePersneling, error)
	GetTipePersnelingByID(ctx context.Context, tipePersnelingID uint8) (entity.TipePersneling, error)
	DeleteTipePersneling(ctx context.Context, tipePersneling uint8) error
	UpdateTipePersneling(ctx context.Context, tipePersneling entity.TipePersneling) error
}

type tipeConnection struct {
	connection *gorm.DB
}

func NewTipeRepository(db *gorm.DB) TipeRepository {
	return &tipeConnection{
		connection: db,
	}
}

func (db *tipeConnection) AddTipeMobil(ctx context.Context, tipeMobil entity.TipeMobil) (entity.TipeMobil, error) {
	uc := db.connection.Create(&tipeMobil)
	if uc.Error != nil {
		return entity.TipeMobil{}, uc.Error
	}
	return tipeMobil, nil
}

func (db *tipeConnection) GetAllTipeMobil(ctx context.Context) ([]entity.TipeMobil, error) {
	var listTipeMobil []entity.TipeMobil
	tx := db.connection.Find(&listTipeMobil)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listTipeMobil, nil
}

func (db *tipeConnection) GetTipeMobilByID(ctx context.Context, tipeMobilID uint8) (entity.TipeMobil, error) {
	var tipeMobil entity.TipeMobil
	tx := db.connection.Where("id = ?", tipeMobilID).Find(&tipeMobil)
	if tx.Error != nil {
		return tipeMobil, tx.Error
	}
	return tipeMobil, nil
}

func (db *tipeConnection) DeleteTipeMobil(ctx context.Context, tipeMobil uint8) error {
	uc := db.connection.Delete(&entity.TipeMobil{}, &tipeMobil)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *tipeConnection) UpdateTipeMobil(ctx context.Context, tipeMobil entity.TipeMobil) error {
	uc := db.connection.Updates(&tipeMobil)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *tipeConnection) AddTipeMesin(ctx context.Context, tipeMesin entity.TipeMesin) (entity.TipeMesin, error) {
	uc := db.connection.Create(&tipeMesin)
	if uc.Error != nil {
		return entity.TipeMesin{}, uc.Error
	}
	return tipeMesin, nil
}

func (db *tipeConnection) GetAllTipeMesin(ctx context.Context) ([]entity.TipeMesin, error) {
	var listTipeMesin []entity.TipeMesin
	tx := db.connection.Find(&listTipeMesin)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listTipeMesin, nil
}

func (db *tipeConnection) GetTipeMesinByID(ctx context.Context, tipeMesinID uint8) (entity.TipeMesin, error) {
	var tipeMesin entity.TipeMesin
	tx := db.connection.Where("id = ?", tipeMesinID).Find(&tipeMesin)
	if tx.Error != nil {
		return tipeMesin, tx.Error
	}
	return tipeMesin, nil
}

func (db *tipeConnection) DeleteTipeMesin(ctx context.Context, tipeMesin uint8) error {
	uc := db.connection.Delete(&entity.TipeMesin{}, &tipeMesin)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *tipeConnection) UpdateTipeMesin(ctx context.Context, tipeMesin entity.TipeMesin) error {
	uc := db.connection.Updates(&tipeMesin)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *tipeConnection) AddTipePersneling(ctx context.Context, tipePersneling entity.TipePersneling) (entity.TipePersneling, error) {
	uc := db.connection.Create(&tipePersneling)
	if uc.Error != nil {
		return entity.TipePersneling{}, uc.Error
	}
	return tipePersneling, nil
}

func (db *tipeConnection) GetAllTipePersneling(ctx context.Context) ([]entity.TipePersneling, error) {
	var listTipePersneling []entity.TipePersneling
	tx := db.connection.Find(&listTipePersneling)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listTipePersneling, nil
}

func (db *tipeConnection) GetTipePersnelingByID(ctx context.Context, tipePersnelingID uint8) (entity.TipePersneling, error) {
	var tipePersneling entity.TipePersneling
	tx := db.connection.Where("id = ?", tipePersnelingID).Find(&tipePersneling)
	if tx.Error != nil {
		return tipePersneling, tx.Error
	}
	return tipePersneling, nil
}

func (db *tipeConnection) DeleteTipePersneling(ctx context.Context, tipePersneling uint8) error {
	uc := db.connection.Delete(&entity.TipePersneling{}, &tipePersneling)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *tipeConnection) UpdateTipePersneling(ctx context.Context, tipePersneling entity.TipePersneling) error {
	uc := db.connection.Updates(&tipePersneling)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
