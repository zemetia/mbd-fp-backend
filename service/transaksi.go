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
	GetTransaksi(ctx context.Context, transaksiID string) (entity.Transaksi, error)
	DeleteTransaksi(ctx context.Context, transaksiID uuid.UUID) error
	UpdateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiUpdateDto) error
	KembaliMobilTransaksi(ctx context.Context, transaksiID uuid.UUID) error
}

type transaksiService struct {
	transaksiRepository  repository.TransaksiRepository
	mobilRepository      repository.MobilRepository
	userRepository       repository.UserRepository
	membershipRepository repository.MembershipRepository
}

func NewTransaksiService(tr repository.TransaksiRepository, mr repository.MobilRepository, ur repository.UserRepository, mr2 repository.MembershipRepository) TransaksiService {
	return &transaksiService{
		transaksiRepository:  tr,
		mobilRepository:      mr,
		userRepository:       ur,
		membershipRepository: mr2,
	}
}

func (us *transaksiService) AddTransaksi(ctx context.Context, transaksiDTO dto.TransaksiCreateDto) (entity.Transaksi, error) {
	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))

	if err != nil {
		return transaksi, err
	}

	mobil, err := us.mobilRepository.FindMobilByID(ctx, transaksi.MobilID)

	if err != nil {
		return transaksi, err
	}

	membership, err := us.membershipRepository.GetMembershipByUserId(ctx, transaksi.UserID)

	if err != nil {
		return transaksi, err
	}

	timediff := float32(transaksi.TglAmbil.Sub(transaksi.TglKembali)/24 + 1)
	transaksi.TotalHarga = timediff * mobil.Price * (1 - (float32(membership.Diskon) / 100))

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

func (us *transaksiService) KembaliMobilTransaksi(ctx context.Context, transaksiID string) error {
	return us.transaksiRepository.UpdateMobilKembali(ctx, transaksiID)
}

func (us *transaksiService) UpdateTransaksi(ctx context.Context, transaksiDTO dto.TransaksiUpdateDto) error {
	transaksi := entity.Transaksi{}
	err := smapping.FillStruct(&transaksi, smapping.MapFields(transaksiDTO))
	if err != nil {
		return err
	}
	return us.transaksiRepository.UpdateTransaksi(ctx, transaksi)
}
