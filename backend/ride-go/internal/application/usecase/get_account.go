package usecase

import (
	"github.com/leonardograselalmeida/fake_uber/internal/application/repository"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
)

type GetAccount struct {
	AccountRepository repository.AccountRepositoryInterface
}

func (g *GetAccount) Execute(accountId string) *entity.Account {
	account, err := g.AccountRepository.GetAccountById(accountId)

	if err != nil {
		return nil
	}

	return account
}
