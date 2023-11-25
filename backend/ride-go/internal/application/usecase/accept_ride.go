package usecase

import (
	"errors"

	"github.com/leonardograselalmeida/fake_uber/internal/application/repository"
)

type AcceptRide struct {
	AccountRepository repository.AccountRepositoryInterface
	RideRepository    repository.RideRepositoryInterface
}

type AcceptRideInput struct {
	DriverId string
	RideId   string
}

func (a *AcceptRide) Execute(input AcceptRideInput) error {
	account, accountError := a.AccountRepository.GetAccountById(input.DriverId)

	if accountError != nil {
		return accountError
	}

	if !account.IsDriver {
		return errors.New("only drivers can accept a ride")
	}

	ride, rideError := a.RideRepository.GetRideById(input.RideId)

	if rideError != nil {
		return rideError
	}

	if ride == nil {
		return errors.New("ride not found")
	}

	ride.Accept(input.DriverId)
	a.RideRepository.UpdateRide(ride)

	return nil
}
