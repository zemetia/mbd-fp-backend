package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/mashingan/smapping"
)

type MobilService interface {
	AddMobil(ctx context.Context, mobilDTO dto.MobilCreateDto) (entity.Mobil, error)
	GetAllMobil(ctx context.Context) ([]entity.Mobil, error)
	GetMobil(ctx context.Context, mobilID string) (entity.Mobil, error)
	DeleteMobil(ctx context.Context, mobilID string) error
	UpdateMobil(ctx context.Context, mobilDTO dto.MobilUpdateDto) error
	GetMobilByLokasiID(ctx context.Context, lokasiID string) ([]entity.Mobil, error)
}

type mobilService struct {
	mobilRepository repository.MobilRepository
	userRepository  repository.UserRepository
	tipeRepository  repository.TipeRepository
}

func NewMobilService(mr repository.MobilRepository, ur repository.UserRepository, tr repository.TipeRepository) MobilService {
	return &mobilService{
		mobilRepository: mr,
		userRepository:  ur,
		tipeRepository:  tr,
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

func (us *mobilService) GetMobil(ctx context.Context, mobilID string) (entity.Mobil, error) {
	return us.mobilRepository.FindMobilByID(ctx, mobilID)
}

func (us *mobilService) GetMobilByLokasiID(ctx context.Context, lokasiID string) ([]entity.Mobil, error) {
	mobilList, err := us.mobilRepository.GetMobilByLokasiID(ctx, lokasiID)
	if err != nil {
		return mobilList, err
	}
	return mobilList, nil
}

func (us *mobilService) DeleteMobil(ctx context.Context, mobilID string) error {
	return us.mobilRepository.DeleteMobil(ctx, mobilID)
}

func (us *mobilService) UpdateMobil(ctx context.Context, mobilDTO dto.MobilUpdateDto) error {
	newMobil := entity.Mobil{}

	err := smapping.FillStruct(&newMobil, smapping.MapFields(mobilDTO))
	if err != nil {
		return err
	}
	return us.mobilRepository.UpdateMobil(ctx, newMobil)
}
