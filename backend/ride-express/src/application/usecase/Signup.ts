import Account from '../../domain/Account';
import type Logger from '../logger/Logger';
import type AccountRepository from '../repository/AccountRepository';

export default class Signup {
  constructor(
    private accountRepository: AccountRepository,
    private logger: Logger
  ) {}

  async execute(input: SignupInput): Promise<SignupOutput> {
    this.logger.log(`signup ${input.name}`);
    const existingAccount = await this.accountRepository.getByEmail(
      input.email
    );
    if (existingAccount) throw new Error('Duplicated account');
    const account = Account.create(
      input.name,
      input.email,
      input.cpf,
      input.carPlate || '',
      !!input.isPassenger,
      !!input.isDriver
    );
    await this.accountRepository.save(account);
    return {
      accountId: account.accountId
    };
  }
}

export type SignupInput = {
  name: string;
  email: string;
  cpf: string;
  carPlate?: string;
  isPassenger?: boolean;
  isDriver?: boolean;
  password: string;
};

type SignupOutput = {
  accountId: string;
};