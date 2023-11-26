package usecase

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/application/logger"
	"github.com/leonardograselalmeida/fake_uber/internal/application/repository"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
)

type Signup struct {
	AccountRepository repository.AccountRepositoryInterface
	Logger            logger.LoggerInterface
}

type SignupInput struct {
	Name        string
	Email       string
	Cpf         string
	CarPlate    string
	IsPassenger bool
	IsDriver    bool
	Password    string
}

type SignupOutput struct {
	AccountId uuid.UUID
}

func (s *Signup) Execute(input SignupInput) (*SignupOutput, error) {
	messageLog := fmt.Sprintf("signup %s", input.Name)
	s.Logger.Log(messageLog)

	existingAccount, existingAccountError := s.AccountRepository.GetAccountByEmail(input.Email)

	if existingAccountError != nil {
		return nil, existingAccountError
	}

	if existingAccount != nil {
		return nil, errors.New("duplicated account")
	}

	account, accountError := entity.CreateAccount(input.Name, input.Email, input.Cpf, input.CarPlate, input.IsPassenger, input.IsDriver)

	if accountError != nil {
		return nil, accountError
	}

	s.AccountRepository.SaveAccount(account)

	output := SignupOutput{
		AccountId: account.AccountId,
	}

	return &output, nil

}
