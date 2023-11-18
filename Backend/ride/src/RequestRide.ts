import crypto from 'crypto';
import type Logger from './Logger';
import type RideDAO from './RideDAO';

export default class RequestRide {
  constructor(
    private rideDAO: RideDAO,
    private logger: Logger
  ) {}

  async execute(input: any) {
    this.logger.log(`requestRide`);
    input.rideId = crypto.randomUUID();
    input.status = 'requested';
    input.date = new Date();
    await this.rideDAO.save(input);
    return {
      rideId: input.rideId
    };
  }
}
