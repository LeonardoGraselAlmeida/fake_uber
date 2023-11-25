package usecase

// import type RideDAO from '../repository/RideRepository';

// export default class StartRide {
//   constructor(private rideDAO: RideDAO) {}

//   async execute(input: StartRideInput): Promise<void> {
//     const ride = await this.rideDAO.getById(input.rideId);
//     if (!ride) throw new Error('Ride not found');
//     ride.start();
//     await this.rideDAO.update(ride);
//   }
// }

// export type StartRideInput = {
//   rideId: string;
// };
