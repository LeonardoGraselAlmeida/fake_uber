import type AccountGateway from '../gateway/AccountGateway';
import type RideRepository from '../repository/RideRepository';

export default class AcceptRide {
  constructor(
    private rideRepository: RideRepository,
    private accountGateway: AccountGateway
  ) {}

  async execute(input: any) {
    const account = await this.accountGateway.getById(input.driverId);
    if (account && !account.isDriver)
      throw new Error('Only drivers can accept a ride');
    const ride = await this.rideRepository.getById(input.rideId);
    if (!ride) throw new Error('Ride not found');
    ride.accept(input.driverId);
    await this.rideRepository.update(ride);
  }
}
