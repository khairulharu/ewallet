package service

import (
	"context"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUser(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {

	return dto.AuthRes{}, nil
}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	return dto.UserData{}, nil
}
