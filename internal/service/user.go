package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/khairulharu/ewallet/domain"
	"github.com/khairulharu/ewallet/dto"
	"github.com/khairulharu/ewallet/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
	emailService    domain.EmailService
}

func NewUser(userRepository domain.UserRepository, cacheRepository domain.CacheRepository, emailService domain.EmailService) domain.UserService {
	return &userService{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
		emailService:    emailService,
	}
}

func (u userService) Authenticate(ctx context.Context, req dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthRes{}, err
	}

	if user == (domain.User{}) {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthRes{}, domain.ErrAuthFailed
	}

	token := util.GenerateRandomString(16)

	userJson, _ := json.Marshal(user)
	_ = u.cacheRepository.Set("user:"+token, userJson)
	return dto.AuthRes{
		Token:  token,
		UserID: user.ID,
	}, nil

}

func (u userService) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := u.cacheRepository.Get("user:" + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	var user domain.User

	_ = json.Unmarshal(data, &user)

	return dto.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		Username: user.Username,
	}, nil
}

func (u userService) Register(ctx context.Context, req dto.UserRegisterReq) (dto.UserRegisterRes, error) {
	exist, err := u.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	if exist != (domain.User{}) {
		return dto.UserRegisterRes{}, domain.ErrUsernameTaken
	}

	hashedPas, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	user := domain.User{
		FullName: req.FullName,
		Phone:    req.Phone,
		Username: req.Username,
		Password: string(hashedPas),
		Email:    req.Email,
	}

	err = u.userRepository.Insert(ctx, &user)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	otpCode := util.GenerateRandomNumber(4)
	referenceId := util.GenerateRandomString(16)

	log.Println("your otp code: ", otpCode)
	// this code has disbale
	// if err := u.emailService.Send(req.Email, "otp code brudaa", otpCode); err != nil {
	// 	return dto.UserRegisterRes{}, err
	// }
	_ = u.cacheRepository.Set("otp:"+referenceId, []byte(otpCode))
	_ = u.cacheRepository.Set("user-ref:"+referenceId, []byte(user.Username))
	return dto.UserRegisterRes{
		ReferenceID: referenceId,
	}, nil
}

func (u userService) ValidateOTP(ctx context.Context, req dto.ValidateOtpReq) error {
	val, err := u.cacheRepository.Get("otp:" + req.ReferenceID)
	if err != nil {
		return domain.ErrOtpInvalid
	}

	otp := string(val)

	if otp != req.OTP {
		return domain.ErrOtpInvalid
	}

	val, err = u.cacheRepository.Get("user-ref:" + req.ReferenceID)
	if err != nil {
		return domain.ErrOtpInvalid
	}

	user, err := u.userRepository.FindByUsername(ctx, string(val))
	if err != nil {
		return err
	}

	user.EmailVerifiedAt = time.Now()

	_ = u.userRepository.Update(ctx, &user)
	return nil
}
