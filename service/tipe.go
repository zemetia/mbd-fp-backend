package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/mashingan/smapping"
)

type TipeService interface {
	AddTipeMobil(ctx context.Context, tipeMobilDTO dto.CreateTipeMobil) (entity.TipeMobil, error)
	GetAllTipeMobil(ctx context.Context) ([]entity.TipeMobil, error)
	DeleteTipeMobil(ctx context.Context, tipeMobil uint8) error
	UpdateTipeMobil(ctx context.Context, tipeMobilDTO dto.UpdateTipeMobil) error

	AddTipeMesin(ctx context.Context, tipeMesinDTO dto.CreateTipeMesin) (entity.TipeMesin, error)
	GetAllTipeMesin(ctx context.Context) ([]entity.TipeMesin, error)
	DeleteTipeMesin(ctx context.Context, tipeMesin uint8) error
	UpdateTipeMesin(ctx context.Context, tipeMesinDTO dto.UpdateTipeMesin) error

	AddTipePersneling(ctx context.Context, tipePersnelingDTO dto.CreateTipePersneling) (entity.TipePersneling, error)
	GetAllTipePersneling(ctx context.Context) ([]entity.TipePersneling, error)
	DeleteTipePersneling(ctx context.Context, tipePersneling uint8) error
	UpdateTipePersneling(ctx context.Context, tipePersnelingDTO dto.UpdateTipePersneling) error
}

type tipeService struct {
	tipeRepository repository.TipeRepository
}

func NewTipeService(tr repository.TipeRepository) TipeService {
	return &tipeService{
		tipeRepository: tr,
	}
}

func (ts *tipeService) AddTipeMobil(ctx context.Context, tipeMobilDTO dto.CreateTipeMobil) (entity.TipeMobil, error) {
	tipeMobil := entity.TipeMobil{}
	err := smapping.FillStruct(&tipeMobil, smapping.MapFields(tipeMobilDTO))
	if err != nil {
		return tipeMobil, err
	}
	return ts.tipeRepository.AddTipeMobil(ctx, tipeMobil)
}

func (ts *tipeService) GetAllTipeMobil(ctx context.Context) ([]entity.TipeMobil, error) {
	return ts.tipeRepository.GetAllTipeMobil(ctx)
}

func (ts *tipeService) DeleteTipeMobil(ctx context.Context, tipeMobil uint8) error {
	return ts.tipeRepository.DeleteTipeMobil(ctx, tipeMobil)
}

func (ts *tipeService) UpdateTipeMobil(ctx context.Context, tipeMobilDTO dto.UpdateTipeMobil) error {
	tipeMobil := entity.TipeMobil{}
	err := smapping.FillStruct(&tipeMobil, smapping.MapFields(tipeMobilDTO))
	if err != nil {
		return err
	}
	return ts.tipeRepository.UpdateTipeMobil(ctx, tipeMobil)
}

func (ts *tipeService) AddTipeMesin(ctx context.Context, tipeMesinDTO dto.CreateTipeMesin) (entity.TipeMesin, error) {
	tipeMesin := entity.TipeMesin{}
	err := smapping.FillStruct(&tipeMesin, smapping.MapFields(tipeMesinDTO))
	if err != nil {
		return tipeMesin, err
	}
	return ts.tipeRepository.AddTipeMesin(ctx, tipeMesin)
}

func (ts *tipeService) GetAllTipeMesin(ctx context.Context) ([]entity.TipeMesin, error) {
	return ts.tipeRepository.GetAllTipeMesin(ctx)
}

func (ts *tipeService) DeleteTipeMesin(ctx context.Context, tipeMesin uint8) error {
	return ts.tipeRepository.DeleteTipeMesin(ctx, tipeMesin)
}

func (ts *tipeService) UpdateTipeMesin(ctx context.Context, tipeMesinDTO dto.UpdateTipeMesin) error {
	tipeMesin := entity.TipeMesin{}
	err := smapping.FillStruct(&tipeMesin, smapping.MapFields(tipeMesinDTO))
	if err != nil {
		return err
	}
	return ts.tipeRepository.UpdateTipeMesin(ctx, tipeMesin)
}

func (ts *tipeService) AddTipePersneling(ctx context.Context, tipePersnelingDTO dto.CreateTipePersneling) (entity.TipePersneling, error) {
	tipePersneling := entity.TipePersneling{}
	err := smapping.FillStruct(&tipePersneling, smapping.MapFields(tipePersnelingDTO))
	if err != nil {
		return tipePersneling, err
	}
	return ts.tipeRepository.AddTipePersneling(ctx, tipePersneling)
}

func (ts *tipeService) GetAllTipePersneling(ctx context.Context) ([]entity.TipePersneling, error) {
	return ts.tipeRepository.GetAllTipePersneling(ctx)
}

func (ts *tipeService) DeleteTipePersneling(ctx context.Context, tipePersneling uint8) error {
	return ts.tipeRepository.DeleteTipePersneling(ctx, tipePersneling)
}

func (ts *tipeService) UpdateTipePersneling(ctx context.Context, tipePersnelingDTO dto.UpdateTipePersneling) error {
	tipePersneling := entity.TipePersneling{}
	err := smapping.FillStruct(&tipePersneling, smapping.MapFields(tipePersnelingDTO))
	if err != nil {
		return err
	}
	return ts.tipeRepository.UpdateTipePersneling(ctx, tipePersneling)
}
