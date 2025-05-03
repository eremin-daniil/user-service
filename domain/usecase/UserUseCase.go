package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"user_service/boundary/dto"
	"user_service/boundary/gateway"
	"user_service/boundary/repository"
	userEntity "user_service/domain/entity/user"
	userPrimitive "user_service/domain/entity/user/primitive"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrProfileNotCreated = errors.New("profile not created")
)

type UserUseCase struct {
	repository                repository.UserRepositoryInterface
	userEventDrivenGateway    gateway.UserEventDrivenGatewayInterface
	profileEventDrivenGateway gateway.ProfileEventDrivenGatewayInterface
}

func NewUserUseCase(
	userEventDrivenGateway gateway.UserEventDrivenGatewayInterface,
	profileEventDrivenGateway gateway.ProfileEventDrivenGatewayInterface,
	repository repository.UserRepositoryInterface,
) *UserUseCase {
	return &UserUseCase{
		userEventDrivenGateway:    userEventDrivenGateway,
		profileEventDrivenGateway: profileEventDrivenGateway,
		repository:                repository,
	}
}

func (uc *UserUseCase) RegisterUser(ctx context.Context, userDTO *dto.UserDTO) (string, error) {
	user, err := userDTO.ToEntity()
	if err != nil {
		return "", err
	}
	_, err = uc.repository.Create(ctx, user)
	if err != nil {
		return "", err
	}
	err = uc.userEventDrivenGateway.SendCreateUserEvent(ctx, dto.UserToDTO(user))
	if err != nil {
		return "", err
	}
	return user.Id().String(), nil
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, userID string) (*dto.UserDTO, error) {
	entity, err := uc.getUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return dto.UserToDTO(entity), nil
}

func (uc *UserUseCase) GetUserByLogin(ctx context.Context, login string) (*dto.UserDTO, error) {
	loginPrimitive, err := userPrimitive.LoginFromString(login)
	if err != nil {
		return nil, err
	}
	entity, err := uc.repository.GetByLogin(ctx, loginPrimitive)
	if err != nil {
		return nil, err
	}
	return dto.UserToDTO(entity), nil
}

func (uc *UserUseCase) CreateProfile(ctx context.Context, userID string, profileDTO *dto.ProfileDTO) (string, error) {
	user, err := uc.getUserByID(ctx, userID)
	if err != nil {
		return "", err
	}
	profile, err := profileDTO.ToEntity()
	if err != nil {
		return "", err
	}
	user.SetProfile(profile)
	err = uc.repository.Update(ctx, user)
	if err != nil {
		return "", err
	}
	err = uc.profileEventDrivenGateway.SendCreateProfileEvent(ctx, dto.ProfileToDTO(profile))
	if err != nil {
		return "", err
	}
	return profile.Id().String(), nil
}

func (uc *UserUseCase) GetProfileByUserID(ctx context.Context, userID string) (*dto.ProfileDTO, error) {
	user, err := uc.getUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	profile := user.Profile()
	if profile == nil {
		return nil, ErrProfileNotCreated
	}
	return dto.ProfileToDTO(profile), nil
}

func (uc *UserUseCase) UpdateProfile(ctx context.Context, userID string, profileDTO *dto.ProfileDTO) error {
	user, err := uc.getUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user.Profile() == nil {
		return ErrProfileNotCreated
	}
	err = user.Profile().ChangeFirstAndLastName(profileDTO.LastName, profileDTO.FirstName)
	if err != nil {
		return err
	}
	err = uc.repository.Update(ctx, user)
	if err != nil {
		return err
	}
	err = uc.profileEventDrivenGateway.SendUpdateProfileEvent(ctx, dto.ProfileToDTO(user.Profile()))
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) getUserByID(ctx context.Context, userID string) (*userEntity.User, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	entityID := userEntity.ID{UUID: id}
	user, err := uc.repository.GetByID(ctx, entityID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
