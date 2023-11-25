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
	RideId      string
	PassengerId string
	driverId    string
	status      string
	Date        time.Time
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
}

func newRide(rideId string, passengerId string, driverId string, status string, date time.Time, fromLat float64, fromLong float64, toLat float64, toLong float64) *Ride {
	ride := Ride{
		RideId:      rideId,
		PassengerId: passengerId,
		driverId:    driverId,
		status:      status,
		Date:        date,
		FromLat:     fromLat,
		FromLong:    fromLong,
		ToLat:       toLat,
		ToLong:      toLong,
	}

	return &ride
}

func CreateRide(passengerId string, fromLat float64, fromLong float64, toLat float64, toLong float64) *Ride {
	rideId := uuid.New().String()
	driverId := ""
	status := StatusRequested
	currentTime := time.Now()
	ride := newRide(rideId, passengerId, driverId, status, currentTime, fromLat, fromLong, toLat, toLong)
	return ride
}

func (ride *Ride) Accept(driverId string) {
	ride.driverId = driverId
	ride.status = StatusAccept
}

func (ride *Ride) Start() {
	ride.status = StatusInProgres
}

func (ride *Ride) GetStatus() string {
	return ride.status
}

func (ride *Ride) GetDriverId() string {
	return ride.driverId
}
