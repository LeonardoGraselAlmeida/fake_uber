package entity

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ride_Valid(t *testing.T) {
	expectPassengerId := uuid.New()
	expectFromLat := 10.000
	expectFromLong := -10.000
	expectToLat := 20.000
	expectToLong := -20.000
	expectDriverId := uuid.Nil
	expectStatus := StatusRequested

	ride := CreateRide(expectPassengerId, expectFromLat, expectFromLong, expectToLat, expectToLong)

	assert.NotNil(t, ride.RideId)
	assert.NotNil(t, ride.Date)
	assert.EqualValues(t, ride.PassengerId, expectPassengerId)
	assert.EqualValues(t, ride.DriverId, expectDriverId)
	assert.EqualValues(t, ride.Status, expectStatus)
	assert.EqualValues(t, ride.FromLat, expectFromLat)
	assert.EqualValues(t, ride.FromLong, expectFromLong)
	assert.EqualValues(t, ride.ToLat, expectToLat)
	assert.EqualValues(t, ride.ToLong, expectToLong)

}

func Test_Ride_When_Call_Accept_Should_Set_DriverId_Change_Status(t *testing.T) {
	expectPassengerId := uuid.New()
	expectFromLat := 10.000
	expectFromLong := -10.000
	expectToLat := 20.000
	expectToLong := -20.000
	expectDriverId := uuid.New()
	expectStatus := StatusAccept

	ride := CreateRide(expectPassengerId, expectFromLat, expectFromLong, expectToLat, expectToLong)
	ride.Accept(expectDriverId)

	assert.Equal(t, ride.DriverId, expectDriverId)
	assert.Equal(t, ride.Status, expectStatus)
}

func Test_Ride_When_Call_Start_Should_Status(t *testing.T) {
	expectPassengerId := uuid.New()
	expectFromLat := 10.000
	expectFromLong := -10.000
	expectToLat := 20.000
	expectToLong := -20.000
	expectStatus := StatusInProgres

	ride := CreateRide(expectPassengerId, expectFromLat, expectFromLong, expectToLat, expectToLong)
	ride.Start()

	assert.Equal(t, ride.Status, expectStatus)
}