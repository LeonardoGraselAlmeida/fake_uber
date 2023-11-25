package usecase

import "github.com/leonardograselalmeida/fake_uber/internal/application/repository"

type AcceptRide struct {
	AccountDAO repository.AccountRepositoryInterface
}

type AcceptRideInput struct {
	DriverId string
	RideId   string
}

func (a *AcceptRide) Execute(input AcceptRideInput) {
	// const account = await this.accountDAO.getById(input.driverId);
	// if (account && !account.isDriver)
	//
	//	throw new Error('Only drivers can accept a ride');
	//
	// const ride = await this.rideDAO.getById(input.rideId);
	// if (!ride) throw new Error('Ride not found');
	// ride.accept(input.driverId);
	// await this.rideDAO.update(ride);
}
