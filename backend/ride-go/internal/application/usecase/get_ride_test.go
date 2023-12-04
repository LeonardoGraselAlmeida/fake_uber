package usecase

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/database"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/logger"
	"github.com/leonardograselalmeida/fake_uber/internal/infra/repository"
	"github.com/stretchr/testify/assert"
)

func Test_Get_Ride_Should_Return_Ride_Valid(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeGetRide(db)
	passengerId := uuid.New()
	expectRide := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	defer db.Close()
	rideRow := createRideRow().
		AddRow(expectRide.RideId, expectRide.PassengerId, expectRide.GetDriverId(), expectRide.GetStatus(), expectRide.FromLat, expectRide.FromLong, expectRide.ToLat, expectRide.ToLong, expectRide.Date)

	setQuerySelectToRide(mock, expectRide.RideId).
		WillReturnRows(rideRow)

	ride, err := sut.Execute(expectRide.RideId)

	if err != nil {
		t.Fatalf("Error in select ride: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.NoError(t, err)
	assert.NotNil(t, ride)
}

func Test_Get_Ride_Should_Return_Ride_Empty(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeGetRide(db)
	rideId := uuid.New()

	defer db.Close()
	rideRow := createRideRow()

	setQuerySelectToRide(mock, rideId).
		WillReturnRows(rideRow)

	ride, err := sut.Execute(rideId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.Nil(t, ride)
	assert.EqualError(t, err, "ride not found")
}

func Test_Get_Ride_Should_Return_SQL_Error(t *testing.T) {
	db, mock := database.NewMockDatabase()
	sut := makeGetRide(db)
	passengerId := uuid.New()
	expectRide := entity.CreateRide(passengerId, 10.000, -10.000, 20.000, -20.000)

	defer db.Close()
	setQuerySelectToRide(mock, expectRide.RideId).
		WillReturnError(errors.New("erro ao selecionar ride no banco de dados"))

	_, err := sut.Execute(expectRide.RideId)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Error in querys: %s", err)
	}

	assert.EqualError(t, err, "erro ao selecionar ride no banco de dados")
}

func makeGetRide(db *sql.DB) GetRide {
	rideRepository := repository.RideRepository{Db: db}
	logger := logger.Logger{}
	sut := GetRide{RideRepository: &rideRepository, Logger: &logger}
	return sut
}
