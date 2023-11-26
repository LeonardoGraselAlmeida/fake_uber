package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
)

type AccountRepository struct {
	Db *sql.DB
}

func (repository *AccountRepository) SaveAccount(account *entity.Account) error {
	_, err := repository.Db.Exec("insert into cccat14.account (account_id, name, email, cpf, car_plate, is_passenger, is_driver) values ($1, $2, $3, $4, $5, $6, $7)",
		account.AccountId, account.Name, account.Email, account.Cpf, account.CarPlate, account.IsPassenger, account.IsDriver)

	if err != nil {
		return err
	}

	return nil
}

func (repository *AccountRepository) GetAccountById(accountId uuid.UUID) (*entity.Account, error) {
	var result entity.Account
	row := repository.Db.QueryRow("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1", accountId)

	if err := row.Scan(&result.AccountId, &result.Name, &result.Email, &result.Cpf, &result.CarPlate, &result.IsPassenger, &result.IsDriver); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	account, errAccount := entity.RestoreAccount(result.AccountId, result.Name, result.Email, result.Cpf, result.CarPlate, result.IsPassenger, result.IsDriver)

	if errAccount != nil {
		return nil, errAccount
	}

	return account, nil
}

func (repository *AccountRepository) GetAccountByEmail(email string) (*entity.Account, error) {
	var result entity.Account
	row := repository.Db.QueryRow("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where email = $1", email)

	if err := row.Scan(&result.AccountId, &result.Name, &result.Email, &result.Cpf, &result.CarPlate, &result.IsPassenger, &result.IsDriver); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	account, errAccount := entity.RestoreAccount(result.AccountId, result.Name, result.Email, result.Cpf, result.CarPlate, result.IsPassenger, result.IsDriver)

	if errAccount != nil {
		return nil, errAccount
	}

	return account, nil
}
