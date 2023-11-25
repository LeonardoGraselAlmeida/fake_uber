package repository

import "github.com/leonardograselalmeida/fake_uber/internal/domain/entity"

type RideRepositoryInterface interface {
	SaveRide(ride *entity.Ride) error
	UpdateRide(ride *entity.Ride) error
	GetRideById(rideId string) (*entity.Ride, error)
	GetActiveRideByPassengerId(passagerId string) (*entity.Ride, error)
	GetAllRide() ([]*entity.Ride, error)
}
