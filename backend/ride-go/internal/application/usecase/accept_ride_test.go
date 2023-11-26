package usecase

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/database"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/repository"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_AcceptRide_When_Ride_And_Driver_Is_Valid(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeSUT(db)
	passengerId := uuid.New()
	driver, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "AAA9999", false, true)
	ride := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	acceptRideInput := AcceptRideInput{
		DriverId: driver.AccountId,
		RideId:   ride.RideId,
	}

	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "nome", "email", "cpf", "car_plate", "is_passenger", "is_driver"}).
		AddRow(driver.AccountId, driver.Name, driver.Email, driver.Cpf, driver.CarPlate, driver.IsPassenger, driver.IsDriver)
	mock.ExpectQuery("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1").
		WithArgs(driver.AccountId).
		WillReturnRows(accountRow)

	rideRow := sqlmock.NewRows([]string{"ride_id", "passenger_id", "driver_id", "status", "from_lat", "from_long", "to_lat", "to_long", "date"}).
		AddRow(ride.RideId, ride.PassengerId, ride.DriverId, ride.Status, ride.FromLat, ride.FromLong, ride.ToLat, ride.ToLong, ride.Date)
	mock.ExpectQuery("select ride_id, passenger_id, driver_id, status, from_lat, from_long, to_lat, to_long, date FROM cccat14.ride where ride_id = $1").
		WithArgs(ride.RideId).
		WillReturnRows(rideRow)

	mock.ExpectExec("update cccat14.ride set status = $1, driver_id = $2 where ride_id = $3").
		WithArgs(entity.StatusAccept, driver.AccountId, ride.RideId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := sut.Execute(acceptRideInput)

	if err != nil {
		t.Fatalf("Error in update ride: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.NoError(t, err)
}

func Test_AcceptRide_When_Ride_Not_Exists_Sould_Return_Error(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeSUT(db)
	passengerId := uuid.New()
	driver, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "AAA9999", false, true)
	ride := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	acceptRideInput := AcceptRideInput{
		DriverId: driver.AccountId,
		RideId:   ride.RideId,
	}

	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "nome", "email", "cpf", "car_plate", "is_passenger", "is_driver"}).
		AddRow(driver.AccountId, driver.Name, driver.Email, driver.Cpf, driver.CarPlate, driver.IsPassenger, driver.IsDriver)
	mock.ExpectQuery("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1").
		WithArgs(driver.AccountId).
		WillReturnRows(accountRow)

	rideRow := sqlmock.NewRows([]string{"ride_id", "passenger_id", "driver_id", "status", "from_lat", "from_long", "to_lat", "to_long", "date"})
	mock.ExpectQuery("select ride_id, passenger_id, driver_id, status, from_lat, from_long, to_lat, to_long, date FROM cccat14.ride where ride_id = $1").
		WithArgs(ride.RideId).
		WillReturnRows(rideRow)

	err := sut.Execute(acceptRideInput)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "ride not found")
}

func Test_AcceptRide_When_Driver_Not_Exists_Should_Return_Error(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeSUT(db)
	passengerId := uuid.New()
	driver, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "AAA9999", false, true)
	ride := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	acceptRideInput := AcceptRideInput{
		DriverId: driver.AccountId,
		RideId:   ride.RideId,
	}

	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "nome", "email", "cpf", "car_plate", "is_passenger", "is_driver"})
	mock.ExpectQuery("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1").
		WithArgs(driver.AccountId).
		WillReturnRows(accountRow)

	err := sut.Execute(acceptRideInput)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "driver not found")
}

func Test_AcceptRide_When_Get_Driver_Should_Return_Error_SqlQuery(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeSUT(db)
	passengerId := uuid.New()
	driver, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "AAA9999", false, true)
	ride := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	acceptRideInput := AcceptRideInput{
		DriverId: driver.AccountId,
		RideId:   ride.RideId,
	}

	defer db.Close()
	mock.ExpectQuery("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1").
		WithArgs(driver.AccountId).
		WillReturnError(errors.New("erro ao consultar account no banco de dados"))

	err := sut.Execute(acceptRideInput)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "erro ao consultar account no banco de dados")
}

func Test_AcceptRide_When_Account_Is_Passeger_Sould_Return_Error(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeSUT(db)
	passengerId := uuid.New()
	driver, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "", true, false)
	ride := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	acceptRideInput := AcceptRideInput{
		DriverId: driver.AccountId,
		RideId:   ride.RideId,
	}

	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "nome", "email", "cpf", "car_plate", "is_passenger", "is_driver"}).
		AddRow(driver.AccountId, driver.Name, driver.Email, driver.Cpf, driver.CarPlate, driver.IsPassenger, driver.IsDriver)
	mock.ExpectQuery("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1").
		WithArgs(driver.AccountId).
		WillReturnRows(accountRow)

	err := sut.Execute(acceptRideInput)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "only drivers can accept a ride")
}

func Test_AcceptRide_When_Get_Ride_Should_Return_Error_SqlQuery(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeSUT(db)
	passengerId := uuid.New()
	driver, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "AAA9999", false, true)
	ride := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	acceptRideInput := AcceptRideInput{
		DriverId: driver.AccountId,
		RideId:   ride.RideId,
	}

	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "nome", "email", "cpf", "car_plate", "is_passenger", "is_driver"}).
		AddRow(driver.AccountId, driver.Name, driver.Email, driver.Cpf, driver.CarPlate, driver.IsPassenger, driver.IsDriver)
	mock.ExpectQuery("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1").
		WithArgs(driver.AccountId).
		WillReturnRows(accountRow)

	mock.ExpectQuery("select ride_id, passenger_id, driver_id, status, from_lat, from_long, to_lat, to_long, date FROM cccat14.ride where ride_id = $1").
		WithArgs(ride.RideId).
		WillReturnError(errors.New("erro ao consultar ride no banco de dados"))

	err := sut.Execute(acceptRideInput)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "erro ao consultar ride no banco de dados")
}

func Test_AcceptRide_When_UpdateRide_Should_Return_Sql_Error(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeSUT(db)
	passengerId := uuid.New()
	driver, _ := entity.CreateAccount("John Driver", "john.driver@gmail.com", "14181694046", "AAA9999", false, true)
	ride := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	acceptRideInput := AcceptRideInput{
		DriverId: driver.AccountId,
		RideId:   ride.RideId,
	}

	defer db.Close()
	accountRow := sqlmock.NewRows([]string{"account_id", "nome", "email", "cpf", "car_plate", "is_passenger", "is_driver"}).
		AddRow(driver.AccountId, driver.Name, driver.Email, driver.Cpf, driver.CarPlate, driver.IsPassenger, driver.IsDriver)
	mock.ExpectQuery("select account_id, name, email, cpf, car_plate, is_passenger, is_driver from cccat14.account where account_id = $1").
		WithArgs(driver.AccountId).
		WillReturnRows(accountRow)

	rideRow := sqlmock.NewRows([]string{"ride_id", "passenger_id", "driver_id", "status", "from_lat", "from_long", "to_lat", "to_long", "date"}).
		AddRow(ride.RideId, ride.PassengerId, ride.DriverId, ride.Status, ride.FromLat, ride.FromLong, ride.ToLat, ride.ToLong, ride.Date)
	mock.ExpectQuery("select ride_id, passenger_id, driver_id, status, from_lat, from_long, to_lat, to_long, date FROM cccat14.ride where ride_id = $1").
		WithArgs(ride.RideId).
		WillReturnRows(rideRow)

	mock.ExpectExec("update cccat14.ride set status = $1, driver_id = $2 where ride_id = $3").
		WithArgs(entity.StatusAccept, driver.AccountId, ride.RideId).
		WillReturnError(errors.New("erro ao atualizar ride no banco de dados"))

	err := sut.Execute(acceptRideInput)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "erro ao atualizar ride no banco de dados")
}

func makeSUT(db *sql.DB) AcceptRide {
	accountRepository := repository.AccountRepository{Db: db}
	rideRepository := repository.RideRepository{Db: db}

	sut := AcceptRide{AccountRepository: &accountRepository, RideRepository: &rideRepository}
	return sut
}
