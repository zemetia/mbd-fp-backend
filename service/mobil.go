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
	GetMobilByLokasiID(ctx context.Context, lokasiID string, lokasiName string, distance float64) ([]dto.MobilGetDto, error)
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

func (us *mobilService) GetMobilByLokasiID(ctx context.Context, lokasiID string, lokasiName string, distance float64) ([]dto.MobilGetDto, error) {
	mobilList, err := us.mobilRepository.GetMobilByLokasiID(ctx, lokasiID)
	var result []dto.MobilGetDto
	if err != nil {
		return result, err
	}
	for _, x := range mobilList {
		user, _ := us.userRepository.FindUserByID(ctx, x.UserID)
		tipeMobil, _ := us.tipeRepository.GetTipeMobilByID(ctx, x.TipeMobilID)
		tipeMesin, _ := us.tipeRepository.GetTipeMesinByID(ctx, x.TipeMesinID)
		tipePersneling, _ := us.tipeRepository.GetTipePersnelingByID(ctx, x.TipePersnelingID)
		result = append(result, dto.MobilGetDto{
			ID:                 x.ID,
			Nama:               x.Nama,
			Price:              x.Price,
			PelatNo:            x.PelatNo,
			KapasitasPenumpang: x.KapasitasPenumpang,
			Status:             x.Status,
			KapasitasMesin:     x.KapasitasMesin,
			PhotoURL:           x.PhotoURL,

			Mitra:            user.Name,
			MitraID:          x.UserID,
			TipeMobil:        tipeMobil.Tipe,
			TipeMobilID:      uint(x.TipeMobilID),
			TipePersneling:   tipePersneling.Tipe,
			TipePersnelingID: uint(x.TipePersnelingID),
			TipeMesin:        tipeMesin.Tipe,
			TipeMesinID:      uint(x.TipePersnelingID),
			Lokasi:           lokasiName,
			LokasiID:         x.LokasiID,
			Distance:         distance,
			Rating:           x.Rating,
		})
	}
	return result, nil
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
