import type GetAccount from '../../application/usecase/GetAccount';
import type Signup from '../../application/usecase/Signup';
import type HttpServer from '../http/HttpServer';

// Interface Adapter
export default class MainController {
  constructor(
    readonly httpServer: HttpServer,
    signup: Signup,
    getAccount: GetAccount
  ) {
    httpServer.register(
      'post',
      '/signup',
      async function (
        params: unknown,
        body: {
          name: string;
          email: string;
          cpf: string;
          carPlate?: string;
          isPassenger?: boolean;
          isDriver?: boolean;
          password: string;
        }
      ) {
        const output = await signup.execute(body);
        return output;
      }
    );

    httpServer.register(
      'get',
      '/accounts/:accountId',
      async function (params: { accountId: string }, body: unknown) {
        const output = await getAccount.execute(params.accountId);
        return output;
      }
    );
  }
}