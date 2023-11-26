package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/leonardograselalmeida/fake_uber/internal/application/logger"
	"github.com/leonardograselalmeida/fake_uber/internal/application/repository"
)

type GetRide struct {
	RideRepository repository.RideRepositoryInterface
	Logger         logger.LoggerInterface
}

type GetRideOutput struct {
	RideId      uuid.UUID
	Status      string
	DriverId    uuid.UUID
	PassengerId uuid.UUID
}

func (g *GetRide) Execute(rideId uuid.UUID) (*GetRideOutput, error) {
	g.Logger.Log("getRide")
	ride, rideError := g.RideRepository.GetRideById(rideId)

	if rideError != nil {
		return nil, rideError
	}

	if ride == nil {
		return nil, errors.New("ride not found")
	}

	output := GetRideOutput{
		RideId:      ride.RideId,
		Status:      ride.Status,
		DriverId:    ride.DriverId,
		PassengerId: ride.PassengerId,
	}

	return &output, nil
}
