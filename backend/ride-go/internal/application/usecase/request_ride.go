package usecase

import (
	"errors"

	"github.com/leonardograselalmeida/fake_uber/internal/application/logger"
	"github.com/leonardograselalmeida/fake_uber/internal/application/repository"
	"github.com/leonardograselalmeida/fake_uber/internal/domain/entity"
)

type RequestRide struct {
	AccountRepository repository.AccountRepositoryInterface
	RideRepository    repository.RideRepositoryInterface
	Logger            logger.LoggerInterface
}

type RequestRideInput struct {
	PassengerId string
	FromLat     float64
	FromLong    float64
	ToLat       float64
	ToLong      float64
}

type RequestRideOutput struct {
	RideId string
}

func (r *RequestRide) Execute(input RequestRideInput) (*RequestRideOutput, error) {
	r.Logger.Log("requestRide")
	account, accountError := r.AccountRepository.GetAccountById(input.PassengerId)

	if accountError != nil {
		return nil, accountError
	}
	if account == nil {
		return nil, errors.New("account does not exit")
	}
	if !account.IsPassenger {
		return nil, errors.New("only passenger can request a ride")
	}

	activeRide, activeRideError := r.RideRepository.GetActiveRideByPassengerId(input.PassengerId)

	if activeRideError != nil {
		return nil, activeRideError
	}

	if activeRide != nil {
		return nil, errors.New("passenger has an active ride")
	}

	ride := entity.CreateRide(input.PassengerId, input.FromLat, input.FromLong, input.ToLat, input.ToLong)

	r.RideRepository.SaveRide(ride)

	output := RequestRideOutput{
		RideId: ride.RideId,
	}

	return &output, nil
}
