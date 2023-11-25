package entity_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ride_Valid(t *testing.T) {
	expectPassengerId := uuid.New().String()
	expectFromLat := 10.000
	expectFromLong := -10.000
	expectToLat := 20.000
	expectToLong := -20.000
	expectDriverId := ""
	expectStatus := entity.StatusRequested

	ride := entity.CreateRide(expectPassengerId, expectFromLat, expectFromLong, expectToLat, expectToLong)

	assert.NotNil(t, ride.RideId)
	assert.NotNil(t, ride.Date)
	assert.EqualValues(t, ride.PassengerId, expectPassengerId)
	assert.EqualValues(t, ride.GetDriverId(), expectDriverId)
	assert.EqualValues(t, ride.GetStatus(), expectStatus)
	assert.EqualValues(t, ride.FromLat, expectFromLat)
	assert.EqualValues(t, ride.FromLong, expectFromLong)
	assert.EqualValues(t, ride.ToLat, expectToLat)
	assert.EqualValues(t, ride.ToLong, expectToLong)

}

func Test_Ride_When_Call_Accept_Should_Set_DriverId_Change_Status(t *testing.T) {
	expectPassengerId := uuid.New().String()
	expectFromLat := 10.000
	expectFromLong := -10.000
	expectToLat := 20.000
	expectToLong := -20.000
	expectDriverId := uuid.New().String()
	expectStatus := entity.StatusAccept

	ride := entity.CreateRide(expectPassengerId, expectFromLat, expectFromLong, expectToLat, expectToLong)
	ride.Accept(expectDriverId)

	assert.Equal(t, ride.GetDriverId(), expectDriverId)
	assert.Equal(t, ride.GetStatus(), expectStatus)
}

func Test_Ride_When_Call_Start_Should_Status(t *testing.T) {
	expectPassengerId := uuid.New().String()
	expectFromLat := 10.000
	expectFromLong := -10.000
	expectToLat := 20.000
	expectToLong := -20.000
	expectStatus := entity.StatusInProgres

	ride := entity.CreateRide(expectPassengerId, expectFromLat, expectFromLong, expectToLat, expectToLong)
	ride.Start()

	assert.Equal(t, ride.GetStatus(), expectStatus)
}