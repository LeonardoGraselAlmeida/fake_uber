package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/application/usecase"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/database"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/logger"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/repository"
)

func main() {
	// Testa a Save Account
	account := getNewAccount()
	accountId := testSaveAccount(account)
	testSaveAccount(account)

	// Testa a Get Account
	if accountId != nil {
		println("Tem accountId")
		testGetAccount(*accountId)
	}

	// Testa a Get All Ride
	testGetAllRide()
}

func getNewAccount() *entity.Account {
	seed := int64(time.Now().UnixNano())
	customSource := rand.NewSource(seed)
	customRand := rand.New(customSource)
	emailRandom := fmt.Sprintf("john.doe.%d@gmail.com", customRand.Intn(9999999))

	account, _ := entity.CreateAccount("John Doe", emailRandom, "97456321558", "", true, false)
	return account
}

func testSaveAccount(account *entity.Account) *uuid.UUID {
	db := database.NewPostgresConnection()

	defer db.Close()
	accountRepository := repository.AccountRepository{
		Db: db,
	}
	looger := logger.Logger{}

	signup := usecase.Signup{
		AccountRepository: &accountRepository,
		Logger:            &looger,
	}

	signupInput := usecase.SignupInput{
		Name:        account.Name,
		Email:       account.Email,
		Cpf:         account.Cpf,
		CarPlate:    account.CarPlate,
		IsPassenger: account.IsPassenger,
		IsDriver:    account.IsDriver,
		Password:    uuid.New().String(),
	}
	signupOutput, err := signup.Execute(signupInput)

	if err != nil {
		println("SignupError: ", err.Error())
		return nil
	}

	println("SignupOutput: ", signupOutput.AccountId.String())

	return &signupOutput.AccountId
}

func testGetAccount(accountId uuid.UUID) {
	db := database.NewPostgresConnection()

	defer db.Close()
	accountRepository := repository.AccountRepository{
		Db: db,
	}

	getAccount := usecase.GetAccount{AccountRepository: &accountRepository}

	account, err := getAccount.Execute(accountId)

	if err != nil {
		println("Error: ", err)
		return
	}

	println("Account: ", account.AccountId.String(), account.Name, account.Email, account.Cpf, account.CarPlate, account.IsPassenger, account.IsDriver)
}

func testGetAllRide() {
	db := database.NewPostgresConnection()

	defer db.Close()
	rideRepository := repository.RideRepository{
		Db: db,
	}

	rides, err := rideRepository.GetAllRide()

	if err != nil {
		println("Error testGetAccount: ", err)
		return
	}

	for _, ride := range rides {
		println("Ride: ", ride.RideId.String(), ride.DriverId.String(), ride.Date.String())
	}
}
