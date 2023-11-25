package entity

import (
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type Account struct {
	AccountId   string
	Name        string
	Email       string
	Cpf         string
	CarPlate    string
	IsPassenger bool
	IsDriver    bool
}

func newAccount(accountId string, name string, email string, cpf string, carPlate string, isPassenger bool, isDriver bool) (*Account, error) {
	if isInvalidName(name) {
		err := errors.New("invalid name")
		return nil, err
	}

	if isInvalidEmail(email) {
		err := errors.New("invalid email")
		return nil, err
	}

	if !ValidateCpf(cpf) {
		err := errors.New("invalid cpf")
		return nil, err
	}

	if isDriver && isInvalidCarPlate(carPlate) {
		err := errors.New("invalid car plate")
		return nil, err
	}

	account := &Account{
		AccountId:   accountId,
		Name:        name,
		Email:       email,
		Cpf:         cpf,
		CarPlate:    carPlate,
		IsPassenger: isPassenger,
		IsDriver:    isDriver,
	}

	return account, nil
}

func CreateAccount(name string, email string, cpf string, carPlate string, isPassenger bool, isDriver bool) (*Account, error) {
	accountId := uuid.New().String()
	account, err := newAccount(accountId, name, email, cpf, carPlate, isPassenger, isDriver)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func RestoreAccount(accountId string, name string, email string, cpf string, carPlate string, isPassenger bool, isDriver bool) (*Account, error) {
	account, err := newAccount(accountId, name, email, cpf, carPlate, isPassenger, isDriver)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func isInvalidName(name string) bool {
	regex := regexp.MustCompile(`[a-zA-Z] [a-zA-Z]`)
	return !regex.MatchString(name)
}

func isInvalidEmail(email string) bool {
	regex := regexp.MustCompile(`^(.+)@(.+)$`)
	return !regex.MatchString(email)
}

func isInvalidCarPlate(carPlate string) bool {
	regex := regexp.MustCompile(`[A-Z]{3}[0-9]{4}`)
	return !regex.MatchString(carPlate)
}
