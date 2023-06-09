package repository

import (
	"context"
	"fp-mbd-amidrive/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransaksiRepository interface {
	AddTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entity.Transaksi, error)
	FindTransaksiByID(ctx context.Context, transaksiID string) (entity.Transaksi, error)
	DeleteTransaksi(ctx context.Context, transaksiID string) error
	UpdateTransaksi(ctx context.Context, transaksi entity.Transaksi) error
	UpdateMobilKembali(ctx context.Context, transaksiID string) error
}

type transaksiConnection struct {
	connection *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiConnection{
		connection: db,
	}
}

func (db *transaksiConnection) AddTransaksi(ctx context.Context, transaksi entity.Transaksi) (entity.Transaksi, error) {
	transaksi.ID = uuid.New().String()

	uc := db.connection.Create(&transaksi)
	if uc.Error != nil {
		return entity.Transaksi{}, uc.Error
	}
	return transaksi, nil
}

func (db *transaksiConnection) GetAllTransaksi(ctx context.Context) ([]entity.Transaksi, error) {
	var listTransaksi []entity.Transaksi
	tx := db.connection.Preload(clause.Associations).Find(&listTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listTransaksi, nil
}

func (db *transaksiConnection) GetTransaksi(ctx context.Context, TransaksiID string) ([]entity.Transaksi, error) {
	var listTransaksi []entity.Transaksi
	tx := db.connection.Preload(clause.Associations).Find(&listTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listTransaksi, nil
}

func (db *transaksiConnection) FindTransaksiByEmail(ctx context.Context, email string) (entity.Transaksi, error) {
	var transaksi entity.Transaksi
	ux := db.connection.Where("email = ?", email).Preload(clause.Associations).Take(&transaksi)
	if ux.Error != nil {
		return transaksi, ux.Error
	}
	return transaksi, nil
}

func (db *transaksiConnection) FindTransaksiByID(ctx context.Context, transaksiID string) (entity.Transaksi, error) {
	var transaksi entity.Transaksi
	ux := db.connection.Where("id = ?", transaksiID).Preload(clause.Associations).Take(&transaksi)
	if ux.Error != nil {
		return transaksi, ux.Error
	}
	return transaksi, nil
}

func (db *transaksiConnection) DeleteTransaksi(ctx context.Context, transaksiID string) error {
	uc := db.connection.Delete(&entity.Transaksi{}, &transaksiID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *transaksiConnection) UpdateTransaksi(ctx context.Context, transaksi entity.Transaksi) error {
	uc := db.connection.Updates(&transaksi)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *transaksiConnection) UpdateMobilKembali(ctx context.Context, transaksiID string) error {
	var transaksi entity.Transaksi

	uc := db.connection.Model(&entity.Transaksi{}).Where("id = ?", transaksiID).Update("tgl_dikembalikan", time.Now())
	if uc.Error != nil {
		return uc.Error
	}

	mc := db.connection.Model(&entity.Mobil{}).Where("id = ?", transaksi.MobilID).Update("status", true)
	if mc.Error != nil {
		return uc.Error
	}
	return nil
}
