import type GetAccount from '../../application/usecase/GetAccount';
import type Signup from '../../application/usecase/Signup';
import { inject } from '../di/Registry';
import type HttpServer from '../http/HttpServer';

// Interface Adapter
export default class MainController {
  @inject('httpServer')
  httpServer?: HttpServer;
  @inject('signup')
  signup?: Signup;
  @inject('getAccount')
  getAccount?: GetAccount;

  constructor() {
    this.httpServer?.register(
      'post',
      '/signup',
      async (params: any, body: any) => {
        const output = await this.signup?.execute(body);
        return output;
      }
    );

    this.httpServer?.register(
      'get',
      '/accounts/:accountId',
      async (params: any, body: any) => {
        const output = await this.getAccount?.execute(params.accountId);
        return output;
      }
    );
  }
}
