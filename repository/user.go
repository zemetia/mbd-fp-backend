package repository

import (
	"context"
	"fp-mbd-amidrive/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, user entity.User) (entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
	FindUserByEmail(ctx context.Context, email string) (entity.User, error)
	FindUserByID(ctx context.Context, userID string) (entity.User, error)
	DeleteUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, user entity.User) error
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) RegisterUser(ctx context.Context, user entity.User) (entity.User, error) {
	user.ID = uuid.New().String()
	uc := db.connection.Create(&user)
	if uc.Error != nil {
		return entity.User{}, uc.Error
	}
	return user, nil
}

func (db *userConnection) GetAllUser(ctx context.Context) ([]entity.User, error) {
	var listUser []entity.User
	tx := db.connection.Preload("Membership").Find(&listUser)
	// tx := db.connection.Raw("SELECT * FROM `ayam_geprek`").Scan(&listUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listUser, nil
}

func (db *userConnection) FindUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	ux := db.connection.Where("email = ?", email).Preload("Membership").Take(&user)
	if ux.Error != nil {
		return user, ux.Error
	}
	return user, nil
}

func (db *userConnection) FindUserByID(ctx context.Context, userID string) (entity.User, error) {
	var user entity.User
	ux := db.connection.Where("id = ?", userID).Preload("Membership").Take(&user)
	if ux.Error != nil {
		return user, ux.Error
	}
	return user, nil
}

func (db *userConnection) DeleteUser(ctx context.Context, userID string) error {
	uc := db.connection.Delete(&entity.User{}, &userID)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func (db *userConnection) UpdateUser(ctx context.Context, user entity.User) error {
	uc := db.connection.Updates(&user)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}
