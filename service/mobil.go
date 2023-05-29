package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type MobilService interface {
	AddMobil(ctx context.Context, mobilDTO dto.MobilCreateDto) (entity.Mobil, error)
	GetAllMobil(ctx context.Context) ([]entity.Mobil, error)
	GetMobil(ctx context.Context, mobilID uuid.UUID) (entity.Mobil, error)
	DeleteMobil(ctx context.Context, mobilID uuid.UUID) error
	UpdateMobil(ctx context.Context, mobilDTO dto.MobilUpdateDto) error
}

type mobilService struct {
	mobilRepository repository.MobilRepository
}

func NewMobilService(ur repository.MobilRepository) MobilService {
	return &mobilService{
		mobilRepository: ur,
	}
}

func (us *mobilService) AddMobil(ctx context.Context, mobilDTO dto.MobilCreateDto) (entity.Mobil, error) {
	mobil := entity.Mobil{}
	err := smapping.FillStruct(&mobil, smapping.MapFields(mobilDTO))
	if err != nil {
		return mobil, err
	}
	return us.mobilRepository.AddMobil(ctx, mobil)
}

func (us *mobilService) GetAllMobil(ctx context.Context) ([]entity.Mobil, error) {
	return us.mobilRepository.GetAllMobil(ctx)
}

func (us *mobilService) GetMobil(ctx context.Context, mobilID uuid.UUID) (entity.Mobil, error) {
	return us.mobilRepository.FindMobilByID(ctx, mobilID)
}

func (us *mobilService) DeleteMobil(ctx context.Context, mobilID uuid.UUID) error {
	return us.mobilRepository.DeleteMobil(ctx, mobilID)
}

func (us *mobilService) UpdateMobil(ctx context.Context, mobilDTO dto.MobilUpdateDto) error {
	mobil := entity.Mobil{}
	err := smapping.FillStruct(&mobil, smapping.MapFields(mobilDTO))
	if err != nil {
		return err
	}
	return us.mobilRepository.UpdateMobil(ctx, mobil)
}
