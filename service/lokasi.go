package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type LokasiService interface {
	AddLokasi(ctx context.Context, lokasiDTO dto.LokasiCreateDto) (entity.Lokasi, error)
	GetAllLokasi(ctx context.Context) ([]entity.Lokasi, error)
	GetLokasi(ctx context.Context, lokasiID uuid.UUID) (entity.Lokasi, error)
	DeleteLokasi(ctx context.Context, lokasiID uuid.UUID) error
	UpdateLokasi(ctx context.Context, lokasiDTO dto.LokasiUpdateDto) error
}

type lokasiService struct {
	lokasiRepository repository.LokasiRepository
}

func NewLokasiService(ur repository.LokasiRepository) LokasiService {
	return &lokasiService{
		lokasiRepository: ur,
	}
}

func (us *lokasiService) AddLokasi(ctx context.Context, lokasiDTO dto.LokasiCreateDto) (entity.Lokasi, error) {
	lokasi := entity.Lokasi{}
	err := smapping.FillStruct(&lokasi, smapping.MapFields(lokasiDTO))
	if err != nil {
		return lokasi, err
	}
	return us.lokasiRepository.AddLokasi(ctx, lokasi)
}

func (us *lokasiService) GetAllLokasi(ctx context.Context) ([]entity.Lokasi, error) {
	return us.lokasiRepository.GetAllLokasi(ctx)
}

func (us *lokasiService) GetLokasi(ctx context.Context, lokasiID uuid.UUID) (entity.Lokasi, error) {
	return us.lokasiRepository.FindLokasiByID(ctx, lokasiID)
}

func (us *lokasiService) DeleteLokasi(ctx context.Context, lokasiID uuid.UUID) error {
	return us.lokasiRepository.DeleteLokasi(ctx, lokasiID)
}

func (us *lokasiService) UpdateLokasi(ctx context.Context, lokasiDTO dto.LokasiUpdateDto) error {
	lokasi := entity.Lokasi{}
	err := smapping.FillStruct(&lokasi, smapping.MapFields(lokasiDTO))
	if err != nil {
		return err
	}
	return us.lokasiRepository.UpdateLokasi(ctx, lokasi)
}
