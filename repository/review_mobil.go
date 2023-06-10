package repository

import (
	"context"
	"fp-mbd-amidrive/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	AddReview(ctx context.Context, review entity.ReviewMobil) (entity.ReviewMobil, error)
	UpdateReview(ctx context.Context, review entity.ReviewMobil) error
	DeleteReview(ctx context.Context, reviewID uuid.UUID) error
	GetReviewById(ctx context.Context, reviewID uuid.UUID) (entity.ReviewMobil, error)
	GetAllMobilReview(ctx context.Context, mobilID uuid.UUID) ([]entity.ReviewMobil, error)
	GetAllUserReview(ctx context.Context, userID uuid.UUID) ([]entity.ReviewMobil, error)
}

type reviewConnection struct {
	connection *gorm.DB
}

func NewReviewhipRepository(db *gorm.DB) ReviewRepository {
	return &reviewConnection{
		connection: db,
	}
}

func (db *reviewConnection) AddReview(ctx context.Context, review entity.ReviewMobil) (entity.ReviewMobil, error) {
	mk := db.connection.Create(&review)
	if mk.Error != nil {
		return entity.ReviewMobil{}, mk.Error
	}
	return review, nil
}

func (db *reviewConnection) UpdateReview(ctx context.Context, review entity.ReviewMobil) error {
	mk := db.connection.Updates(&review)
	if mk.Error != nil {
		return mk.Error
	}
	return nil
}

func (db *reviewConnection) DeleteReview(ctx context.Context, reviewID uuid.UUID) error {
	mk := db.connection.Delete(&entity.Membership{}, &reviewID)
	if mk.Error != nil {
		return mk.Error
	}
	return nil
}

func (db *reviewConnection) GetReviewById(ctx context.Context, reviewID uuid.UUID) (entity.ReviewMobil, error) {
	var review entity.ReviewMobil
	mk := db.connection.Where("id = ?", reviewID).Take(&review)
	if mk.Error != nil {
		return review, mk.Error
	}
	return review, nil
}

func (db *reviewConnection) GetAllMobilReview(ctx context.Context, mobilID uuid.UUID) ([]entity.ReviewMobil, error) {
	var listReview []entity.ReviewMobil
	mk := db.connection.Where("mobil_id = ?", mobilID).Find(&listReview)
	if mk.Error != nil {
		return nil, mk.Error
	}
	return listReview, nil
}

func (db *reviewConnection) GetAllUserReview(ctx context.Context, userID uuid.UUID) ([]entity.ReviewMobil, error) {
	var review []entity.ReviewMobil
	mk := db.connection.Where("user_id = ?", userID).Take(&review)
	if mk.Error != nil {
		return review, mk.Error
	}
	return review, nil
}
