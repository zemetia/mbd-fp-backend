package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/mashingan/smapping"
)

type MetodePembayaranService interface {
	AddMetodePembayaran(ctx context.Context, metodePembayaranDTO dto.MetodePembayaranCreateDto) (entity.MetodePembayaran, error)
	GetAllMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error)
	GetMetodePembayaranByID(ctx context.Context, metodePembayaranID uint8) (entity.MetodePembayaran, error)
	DeleteMetodePembayaran(ctx context.Context, metodePembayaranID uint8) error
	UpdateMetodePembayaran(ctx context.Context, metodePembayaranDTO dto.MetodePembayaranUpdateDto) error
}

type metodePembayaranService struct {
	metodePembayaranRepository repository.MetodePembayaranRepository
}

func NewMetodePembayaranService(ur repository.MetodePembayaranRepository) MetodePembayaranService {
	return &metodePembayaranService{
		metodePembayaranRepository: ur,
	}
}

func (us *metodePembayaranService) AddMetodePembayaran(ctx context.Context, metodePembayaranDTO dto.MetodePembayaranCreateDto) (entity.MetodePembayaran, error) {
	metodePembayaran := entity.MetodePembayaran{}
	err := smapping.FillStruct(&metodePembayaran, smapping.MapFields(metodePembayaranDTO))
	if err != nil {
		return metodePembayaran, err
	}
	return us.metodePembayaranRepository.AddMetodePembayaran(ctx, metodePembayaran)
}

func (us *metodePembayaranService) GetAllMetodePembayaran(ctx context.Context) ([]entity.MetodePembayaran, error) {
	return us.metodePembayaranRepository.GetAllMetodePembayaran(ctx)
}

func (us *metodePembayaranService) GetMetodePembayaranByID(ctx context.Context, metodePembayaranID uint8) (entity.MetodePembayaran, error) {
	return us.metodePembayaranRepository.GetMetodePembayaranByID(ctx, metodePembayaranID)
}

func (us *metodePembayaranService) DeleteMetodePembayaran(ctx context.Context, metodePembayaranID uint8) error {
	return us.metodePembayaranRepository.DeleteMetodePembayaran(ctx, metodePembayaranID)
}

func (us *metodePembayaranService) UpdateMetodePembayaran(ctx context.Context, metodePembayaranDTO dto.MetodePembayaranUpdateDto) error {
	metodePembayaran := entity.MetodePembayaran{}
	err := smapping.FillStruct(&metodePembayaran, smapping.MapFields(metodePembayaranDTO))
	if err != nil {
		return err
	}
	return us.metodePembayaranRepository.UpdateMetodePembayaran(ctx, metodePembayaran)
}
