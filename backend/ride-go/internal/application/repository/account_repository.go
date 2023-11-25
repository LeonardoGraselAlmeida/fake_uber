package repository

import "github.com/leonardograselalmeida/fake_uber/internal/domain/entity"

type AccountRepositoryInterface interface {
	SaveAccount(account *entity.Account) error
	GetAccountById(accountId string) (*entity.Account, error)
	GetAccountByEmail(email string) (*entity.Account, error)
}
