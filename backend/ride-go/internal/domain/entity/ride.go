package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	StatusRequested = "requested"
	StatusInProgres = "in_progress"
	StatusAccept    = "accept"
)

type Ride struct {
	RideId      uuid.UUID
	PassengerId uuid.UUID
	DriverId    uuid.UUID
	Status      string
	Date        time.Time
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
}

func newRide(rideId uuid.UUID, passengerId uuid.UUID, driverId uuid.UUID, status string, date time.Time, fromLat float64, fromLong float64, toLat float64, toLong float64) *Ride {
	ride := Ride{
		RideId:      rideId,
		PassengerId: passengerId,
		DriverId:    driverId,
		Status:      status,
		Date:        date,
		FromLat:     fromLat,
		FromLong:    fromLong,
		ToLat:       toLat,
		ToLong:      toLong,
	}

	return &ride
}

func CreateRide(passengerId uuid.UUID, fromLat float64, fromLong float64, toLat float64, toLong float64) *Ride {
	var driverId uuid.UUID
	rideId := uuid.New()
	status := StatusRequested
	currentTime := time.Now()
	ride := newRide(rideId, passengerId, driverId, status, currentTime, fromLat, fromLong, toLat, toLong)
	return ride
}

func RestoreRide(rideId uuid.UUID, passengerId uuid.UUID, driverId uuid.UUID, status string, date time.Time, fromLat float64, fromLong float64, toLat float64, toLong float64) *Ride {
	ride := newRide(rideId, passengerId, driverId, status, date, fromLat, fromLong, toLat, toLong)
	return ride
}

func (ride *Ride) Accept(driverId uuid.UUID) {
	ride.DriverId = driverId
	ride.Status = StatusAccept
}

func (ride *Ride) Start() {
	ride.Status = StatusInProgres
}
