package repository

import (
	"context"
	"fp-mbd-amidrive/entity"

	"gorm.io/gorm"
)

type MembershipRepository interface {
	GetAllMembership(ctx context.Context) ([]entity.Membership, error)
	AddMembership(ctx context.Context, membership entity.Membership) (entity.Membership, error)
	UpdateMembership(ctx context.Context, membership entity.Membership) error
	DeleteMembership(ctx context.Context, membershipID string) error
	GetMembershipByUserId(ctx context.Context, userID string) (entity.Membership, error)
	GetMembershipById(ctx context.Context, membershipID string) (entity.Membership, error)
}

type membershipConnection struct {
	connection *gorm.DB
}

func NewMembershipRepository(db *gorm.DB) MembershipRepository {
	return &membershipConnection{
		connection: db,
	}
}

func (db *membershipConnection) GetAllMembership(ctx context.Context) ([]entity.Membership, error) {
	var listMembership []entity.Membership
	mk := db.connection.Find(&listMembership)
	if mk.Error != nil {
		return nil, mk.Error
	}
	return listMembership, nil
}

func (db *membershipConnection) AddMembership(ctx context.Context, membership entity.Membership) (entity.Membership, error) {
	mk := db.connection.Create(&membership)
	if mk.Error != nil {
		return entity.Membership{}, mk.Error
	}
	return membership, nil
}

func (db *membershipConnection) UpdateMembership(ctx context.Context, membership entity.Membership) error {
	mk := db.connection.Updates(&membership)
	if mk.Error != nil {
		return mk.Error
	}
	return nil
}

func (db *membershipConnection) DeleteMembership(ctx context.Context, membershipID string) error {
	mk := db.connection.Delete(&entity.Membership{}, &membershipID)
	if mk.Error != nil {
		return mk.Error
	}
	return nil
}

func (db *membershipConnection) GetMembershipByUserId(ctx context.Context, userID string) (entity.Membership, error) {
	var user entity.User
	var membership entity.Membership
	mk := db.connection.Where("id = ?", userID).Take(&user)
	if mk.Error != nil {
		return membership, mk.Error
	}
	mk = db.connection.Where("tier = ?", user.MembershipTier).Take(&membership)
	if mk.Error != nil {
		return membership, mk.Error
	}
	return membership, nil
}

func (db *membershipConnection) GetMembershipById(ctx context.Context, membershipID string) (entity.Membership, error) {
	var membership entity.Membership
	mk := db.connection.Where("tier = ?", membershipID).Take(&membership)
	if mk.Error != nil {
		return membership, mk.Error
	}
	return membership, nil
}
