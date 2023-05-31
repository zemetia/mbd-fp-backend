package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type TransaksiService interface {
	AddTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDto) (entity.Transaksi, error)
	GetAllTransaksi(ctx context.Context) ([]entity.Transaksi, error)
	GetTransaksi(ctx context.Context, transaksiID uuid.UUID) (entity.Transaksi, error)
	DeleteTransaksi(ctx context.Context, transaksiID uuid.UUID) error
	UpdateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiUpdateDto) error
}

type transaksiService struct {
	transaksiRepository repository.TransaksiRepository
}

func NewTransaksiService(ur repository.TransaksiRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository: ur,
	}
}

func (us *transaksiService) AddTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDto) (entity.Transaksi, error) {
	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))
	if err != nil {
		return transaksi, err
	}
	return us.transaksiRepository.AddTransaksi(ctx, transaksi)
}

func (us *transaksiService) GetAllTransaksi(ctx context.Context) ([]entity.Transaksi, error) {
	return us.transaksiRepository.GetAllTransaksi(ctx)
}

func (us *transaksiService) GetTransaksi(ctx context.Context, transaksiID uuid.UUID) (entity.Transaksi, error) {
	return us.transaksiRepository.FindTransaksiByID(ctx, transaksiID)
}

func (us *transaksiService) DeleteTransaksi(ctx context.Context, transaksiID uuid.UUID) error {
	return us.transaksiRepository.DeleteTransaksi(ctx, transaksiID)
}

func (us *transaksiService) UpdateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiUpdateDto) error {
	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))
	if err != nil {
		return err
	}
	return us.transaksiRepository.UpdateTransaksi(ctx, transaksi)
}
