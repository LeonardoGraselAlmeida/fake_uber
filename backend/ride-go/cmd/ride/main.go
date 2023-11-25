package main

import (
	"github.com/leonardograselalmeida/fake_uber/internal/application/usecase"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/database"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/repository"
)

func main() {
	db := database.NewPostgresConnection()

	defer db.Close()
	accountRepository := repository.AccountRepository{
		Db: db,
	}

	getAccount := usecase.GetAccount{AccountRepository: &accountRepository}

	account, err := getAccount.Execute("5ff30d7d-eb76-4d65-96ca-ae94221139f6")

	if err != nil {
		println("Error: %s", err)
	}

	println("Account: ", account.AccountId, account.Name, account.Email, account.Cpf, account.CarPlate, account.IsPassenger, account.IsDriver)
}
