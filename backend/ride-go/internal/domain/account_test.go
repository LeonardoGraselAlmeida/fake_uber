package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Passenger_Account_Valid(t *testing.T) {
	expectName := "John Doe"
	expectEmail := "john.doe@gmail.com"
	expectCpf := "97456321558"
	expectCarPlate := ""
	expectIsPassenger := true
	expectIsDriver := false
	account, err := CreateAccount(expectName, expectEmail, expectCpf, expectCarPlate, expectIsPassenger, expectIsDriver)

	assert.NoError(t, err)
	assert.NotNil(t, account.AccountId)
	assert.EqualValues(t, account.Name, expectName)
	assert.EqualValues(t, account.Email, expectEmail)
	assert.EqualValues(t, account.Cpf, expectCpf)
	assert.EqualValues(t, account.CarPlate, expectCarPlate)
	assert.EqualValues(t, account.IsPassenger, expectIsPassenger)
	assert.EqualValues(t, account.IsDriver, expectIsDriver)
}

func Test_Create_Driver_Account_Valid(t *testing.T) {
	expectName := "John Doe"
	expectEmail := "john.doe@gmail.com"
	expectCpf := "97456321558"
	expectCarPlate := "AAA9999"
	expectIsPassenger := false
	expectIsDriver := true
	account, err := CreateAccount(expectName, expectEmail, expectCpf, expectCarPlate, expectIsPassenger, expectIsDriver)

	assert.NoError(t, err)
	assert.NotNil(t, account.AccountId)
	assert.EqualValues(t, account.Name, expectName)
	assert.EqualValues(t, account.Email, expectEmail)
	assert.EqualValues(t, account.Cpf, expectCpf)
	assert.EqualValues(t, account.CarPlate, expectCarPlate)
	assert.EqualValues(t, account.IsPassenger, expectIsPassenger)
	assert.EqualValues(t, account.IsDriver, expectIsDriver)
}

func Test_Restore_Passenger_Account_Valid(t *testing.T) {
	expectAccountId := uuid.New().String()
	expectName := "John Doe"
	expectEmail := "john.doe@gmail.com"
	expectCpf := "97456321558"
	expectCarPlate := ""
	expectIsPassenger := true
	expectIsDriver := false
	account, err := RestoreAccount(expectAccountId, expectName, expectEmail, expectCpf, expectCarPlate, expectIsPassenger, expectIsDriver)

	assert.NoError(t, err)
	assert.EqualValues(t, account.AccountId, expectAccountId)
	assert.EqualValues(t, account.Name, expectName)
	assert.EqualValues(t, account.Email, expectEmail)
	assert.EqualValues(t, account.Cpf, expectCpf)
	assert.EqualValues(t, account.CarPlate, expectCarPlate)
	assert.EqualValues(t, account.IsPassenger, expectIsPassenger)
	assert.EqualValues(t, account.IsDriver, expectIsDriver)
}

func Test_Restore_Driver_Account_Valid(t *testing.T) {
	expectAccountId := uuid.New().String()
	expectName := "John Doe"
	expectEmail := "john.doe@gmail.com"
	expectCpf := "97456321558"
	expectCarPlate := "AAA9999"
	expectIsPassenger := false
	expectIsDriver := true
	account, err := RestoreAccount(expectAccountId, expectName, expectEmail, expectCpf, expectCarPlate, expectIsPassenger, expectIsDriver)

	assert.NoError(t, err)
	assert.EqualValues(t, account.AccountId, expectAccountId)
	assert.EqualValues(t, account.Name, expectName)
	assert.EqualValues(t, account.Email, expectEmail)
	assert.EqualValues(t, account.Cpf, expectCpf)
	assert.EqualValues(t, account.CarPlate, expectCarPlate)
	assert.EqualValues(t, account.IsPassenger, expectIsPassenger)
	assert.EqualValues(t, account.IsDriver, expectIsDriver)
}

func Test_Create_Account_With_Name_Invalid(t *testing.T) {
	expectedErrorMessage := "invalid name"

	_, err := CreateAccount("John", "john.doe@gmail.com", "97456321558", "", true, false)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErrorMessage)
}

func Test_Create_Account_With_Email_Invalid(t *testing.T) {
	expectedErrorMessage := "invalid email"

	_, err := CreateAccount("John Doe", "john@", "97456321558", "", true, false)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErrorMessage)
}

func Test_Create_Account_With_Cpf_Invalid(t *testing.T) {
	expectedErrorMessage := "invalid cpf"

	_, err := CreateAccount("John Doe", "john.doe@gmail.com", "97456321559", "", true, false)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErrorMessage)
}

func Test_Create_Account_With_CarPlate_Invalid(t *testing.T) {
	expectedErrorMessage := "invalid car plate"
	expectAccountId := uuid.New().String()

	_, err := RestoreAccount(expectAccountId, "John Doe", "john.doe@gmail.com", "97456321558", "AAA", false, true)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErrorMessage)
}
