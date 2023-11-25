package usecase

import (
	"errors"

	"github.com/leonardograselalmeida/fake_uber/internal/application/repository"
)

type StartRide struct {
	RideRepository repository.RideRepositoryInterface
}

type StartRideInput struct {
	RideId string
}

func (s *StartRide) Execute(input StartRideInput) error {
	ride, rideError := s.RideRepository.GetRideById(input.RideId)
	if rideError != nil {
		return rideError
	}
	if ride == nil {
		return errors.New("ride not found")
	}
	ride.Start()
	s.RideRepository.UpdateRide(ride)
	return nil
}
