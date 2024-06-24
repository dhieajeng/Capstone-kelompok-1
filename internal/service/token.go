package service

import (
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/pkg/scheduler"
	"github.com/labstack/echo/v4"
)

type tokenService struct {
	otpRepository   repository.OneTimePasswordRepository
	tokenRepository repository.TokenRepository
	userRepository  repository.UserRepository
	scheduler       scheduler.Scheduler
}

func (s *tokenService) GenerateTokenRegistration(c echo.Context, otp, email string) (*entity.Token, error) {
	otpData, err := s.otpRepository.FindOneByCodeAndEmail(c.Request().Context(), email, otp)
	if err != nil || otpData == nil {
		return nil, err
	}

	err = s.otpRepository.Delete(c.Request().Context(), otpData.ID)
	if err != nil {
		return nil, err
	}

	token := entity.NewToken(email, entity.Register)
	token, err = s.tokenRepository.Create(c.Request().Context(), token)
	if err != nil {

		return nil, err
	}
	return token, nil
}

func (s *tokenService) GenerateTokenForgotPassword(c echo.Context, email string) (*entity.Token, error) {
	user, err := s.userRepository.FindByEmail(c.Request().Context(), email)
	if err != nil {
		return nil, err
	}

	token := entity.NewToken(user.Email, entity.ForgotPassword)
	token, err = s.tokenRepository.Create(c.Request().Context(), token)
	if err != nil {
		return nil, err
	}

	schema := "http://"
	if c.Request().TLS != nil {
		schema = "https://"
	}
	link := fmt.Sprintf("%s%sapi/auth/reset-password?token=%v", schema, c.Request().Host, token.ID)

	s.scheduler.SendToken(email, link, token.ID)
	return token, nil
}

type TokenService interface {
	GenerateTokenRegistration(c echo.Context, otp, email string) (*entity.Token, error)
	GenerateTokenForgotPassword(c echo.Context, email string) (*entity.Token, error)
}

func NewTokenService(
	otpRepository repository.OneTimePasswordRepository,
	tokenRepository repository.TokenRepository,
	userRepository repository.UserRepository,
	scheduler scheduler.Scheduler) TokenService {
	return &tokenService{
		otpRepository:   otpRepository,
		tokenRepository: tokenRepository,
		userRepository:  userRepository,
		scheduler:       scheduler,
	}
}
