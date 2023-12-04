package usecase

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/database"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Account_Should_Return_Account_Valid(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeGetAccount(db)
	expectAccount, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "", true, false)

	defer db.Close()
	accountRow := createAccountRow().
		AddRow(expectAccount.AccountId, expectAccount.Name, expectAccount.Email, expectAccount.Cpf, expectAccount.CarPlate, expectAccount.IsPassenger, expectAccount.IsDriver)

	setQuerySelectToAccount(mock, expectAccount.AccountId).
		WillReturnRows(accountRow)

	account, err := sut.Execute(expectAccount.AccountId)

	if err != nil {
		t.Fatalf("Error in select account: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, account)
}

func Test_Get_Account_Should_Return_SQL_Error(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeGetAccount(db)
	expectAccount, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "", true, false)

	defer db.Close()
	setQuerySelectToAccount(mock, expectAccount.AccountId).
		WillReturnError(errors.New("erro ao selecionar account no banco de dados"))

	_, err := sut.Execute(expectAccount.AccountId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "erro ao selecionar account no banco de dados")
}

func makeGetAccount(db *sql.DB) GetAccount {
	accountRepository := repository.AccountRepository{Db: db}
	sut := GetAccount{AccountRepository: &accountRepository}
	return sut
}
