package service

import (
	"context"
	"fp-mbd-amidrive/dto"
	"fp-mbd-amidrive/entity"
	"fp-mbd-amidrive/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type MembershipService interface {
	GetAllMembership(ctx context.Context) ([]entity.Membership, error)
	AddMembership(ctx context.Context, membershipDTO dto.MembershipCreateDto) (entity.Membership, error)
	UpdateMembership(ctx context.Context, membershipDTO dto.MembershipUpdateDto) error
	DeleteMembership(ctx context.Context, membershipID string) error
	GetMembershipByUserId(ctx context.Context, userID uuid.UUID) (entity.Membership, error)
	GetMembershipById(ctx context.Context, membershipID string) (entity.Membership, error)
}

type membershipService struct {
	membershipRepository repository.MembershipRepository
}

func NewMembershipService(ur repository.MembershipRepository) MembershipService {
	return &membershipService{
		membershipRepository: ur,
	}
}

func (ur *membershipService) GetAllMembership(ctx context.Context) ([]entity.Membership, error) {
	return ur.membershipRepository.GetAllMembership(ctx)
}

func (ur *membershipService) AddMembership(ctx context.Context, membershipDTO dto.MembershipCreateDto) (entity.Membership, error) {
	membership := entity.Membership{}
	err := smapping.FillStruct(&membership, smapping.MapFields(membershipDTO))
	if err != nil {
		return membership, err
	}
	return ur.membershipRepository.AddMembership(ctx, membership)
}

func (ur *membershipService) UpdateMembership(ctx context.Context, membershipDTO dto.MembershipUpdateDto) error {
	membership := entity.Membership{}
	err := smapping.FillStruct(&membership, smapping.MapFields(membershipDTO))
	if err != nil {
		return err
	}
	return ur.membershipRepository.UpdateMembership(ctx, membership)
}

func (ur *membershipService) DeleteMembership(ctx context.Context, membershipID string) error {
	return ur.membershipRepository.DeleteMembership(ctx, membershipID)
}

func (ur *membershipService) GetMembershipByUserId(ctx context.Context, userID uuid.UUID) (entity.Membership, error) {
	return ur.membershipRepository.GetMembershipByUserId(ctx, userID)
}

func (ur *membershipService) GetMembershipById(ctx context.Context, membershipID string) (entity.Membership, error) {
	return ur.membershipRepository.GetMembershipById(ctx, membershipID)
}
