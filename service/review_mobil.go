package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/mashingan/smapping"
)

type ReviewService interface {
	AddReview(ctx context.Context, reviewDTO dto.ReviewCreateDto) (entity.ReviewMobil, error)
	DeleteReview(ctx context.Context, reviewID string) error
	UpdateReview(ctx context.Context, reviewDTO dto.ReviewUpdateDto) error
	GetReviewById(ctx context.Context, reviewID string) (entity.ReviewMobil, error)
	GetAllMobilReview(ctx context.Context, mobilID string) ([]entity.ReviewMobil, error)
	GetAllUserReview(ctx context.Context, mobilID string) ([]entity.ReviewMobil, error)
}

type reviewService struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewService(rr repository.ReviewRepository) ReviewService {
	return &reviewService{
		reviewRepository: rr,
	}
}

func (rs *reviewService) AddReview(ctx context.Context, reviewDTO dto.ReviewCreateDto) (entity.ReviewMobil, error) {
	review := entity.ReviewMobil{}
	err := smapping.FillStruct(&review, smapping.MapFields(reviewDTO))
	if err != nil {
		return review, err
	}
	return rs.reviewRepository.AddReview(ctx, review)
}

func (rs *reviewService) GetReviewById(ctx context.Context, reviewID string) (entity.ReviewMobil, error) {
	return rs.reviewRepository.GetReviewById(ctx, reviewID)
}

func (rs *reviewService) GetAllMobilReview(ctx context.Context, mobilID string) ([]entity.ReviewMobil, error) {
	return rs.reviewRepository.GetAllMobilReview(ctx, mobilID)
}

func (rs *reviewService) GetAllUserReview(ctx context.Context, mobilID string) ([]entity.ReviewMobil, error) {
	return rs.reviewRepository.GetAllUserReview(ctx, mobilID)
}

func (rs *reviewService) DeleteReview(ctx context.Context, reviewID string) error {
	return rs.reviewRepository.DeleteReview(ctx, reviewID)
}

func (rs *reviewService) UpdateReview(ctx context.Context, reviewDTO dto.ReviewUpdateDto) error {
	review := entity.ReviewMobil{}
	err := smapping.FillStruct(&review, smapping.MapFields(reviewDTO))
	if err != nil {
		return err
	}
	return rs.reviewRepository.UpdateReview(ctx, review)
}
