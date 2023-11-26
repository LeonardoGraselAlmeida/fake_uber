package usecase

import (
	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/application/repository"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
)

type GetAccount struct {
	AccountRepository repository.AccountRepositoryInterface
}

func (g *GetAccount) Execute(accountId uuid.UUID) (*entity.Account, error) {
	account, err := g.AccountRepository.GetAccountById(accountId)

	if err != nil {
		return nil, err
	}

	return account, nil
}
