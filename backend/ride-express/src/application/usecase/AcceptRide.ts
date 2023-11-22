import type AccountDAO from '../repository/AccountRepository';
import type RideDAO from '../repository/RideRepository';

export default class AcceptRide {
  constructor(
    private rideDAO: RideDAO,
    private accountDAO: AccountDAO
  ) {}

  async execute(input: AcceptRideInput): Promise<void> {
    const account = await this.accountDAO.getById(input.driverId);
    if (account && !account.isDriver)
      throw new Error('Only drivers can accept a ride');
    const ride = await this.rideDAO.getById(input.rideId);
    if (!ride) throw new Error('Ride not found');
    ride.accept(input.driverId);
    await this.rideDAO.update(ride);
  }
}

export type AcceptRideInput = {
  driverId: string;
  rideId: string;
};
