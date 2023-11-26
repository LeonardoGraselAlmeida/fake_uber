package repository

import (
	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
)

type RideRepositoryInterface interface {
	SaveRide(ride *entity.Ride) error
	UpdateRide(ride *entity.Ride) error
	GetRideById(rideId uuid.UUID) (*entity.Ride, error)
	GetActiveRideByPassengerId(passengerId uuid.UUID) (*entity.Ride, error)
	GetAllRide() ([]*entity.Ride, error)
}
