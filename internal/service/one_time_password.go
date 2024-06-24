package service

import (
	"context"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/pkg/scheduler"
)

type oneTimePasswordService struct {
	otpRepository repository.OneTimePasswordRepository
	scheduler     scheduler.Scheduler
}

func (s *oneTimePasswordService) codeGenerator() (code string) {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	c := make([]string, 8)
	for i := range c {
		numOrAlpha := rand.Intn(2)
		if numOrAlpha == 0 {
			c[i] = strconv.Itoa(randomizer.Intn(10))
		} else {
			c[i] = string(letters[randomizer.Intn(len(letters))])
		}

		code = strings.Join(c, "")
	}
	return
}

func (s *oneTimePasswordService) GenerateForRegister(c context.Context, email string) (*entity.OneTimePassword, error) {
	otp := entity.NewOneTimePassword(s.codeGenerator(), email)
	otp, err := s.otpRepository.Create(c, otp)
	if err != nil {
		return nil, err
	}

	s.scheduler.SendOTP(email, otp.OTPCode)
	return otp, nil
}

func (s *oneTimePasswordService) FindOneByCodeAndEmail(c context.Context, email, code string) (*entity.OneTimePassword, error) {
	otp, err := s.otpRepository.FindOneByCodeAndEmail(c, email, code)
	if err != nil {
		return nil, err
	}
	return otp, nil
}

type OneTimePasswordService interface {
	GenerateForRegister(c context.Context, email string) (*entity.OneTimePassword, error)
	FindOneByCodeAndEmail(c context.Context, email, code string) (*entity.OneTimePassword, error)
}

func NewOneTimePasswordService(otpRepository repository.OneTimePasswordRepository, scheduler scheduler.Scheduler) OneTimePasswordService {
	return &oneTimePasswordService{
		otpRepository: otpRepository,
		scheduler:     scheduler,
	}
}
