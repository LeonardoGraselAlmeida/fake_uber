import dotenv from 'dotenv';
import GetAccount from './application/usecase/GetAccount';
import Signup from './application/usecase/Signup';
import MainController from './infra/controller/MainController';
import PgPromiseAdapter from './infra/database/PgPromiseAdapter';
import ExpressAdapter from './infra/http/ExpressAdapter';
import LoggerConsole from './infra/logger/LoggerConsole';
import AccountRepositoryDatabase from './infra/repository/AccountRepositoryDatabase';

dotenv.config();
const portDefaultToListen = 3000;
const portToListen: number = parseInt(process.env.API_PORT || '');

// composition root ou entry point
// criar o grafo de dependências utilizado no projeto

// framework and driver and library
const httpServer = new ExpressAdapter();
const databaseConnection = new PgPromiseAdapter();

// interface adapter
const accountRepository = new AccountRepositoryDatabase(databaseConnection);
const logger = new LoggerConsole();

// use case
const signup = new Signup(accountRepository, logger);
const getAccount = new GetAccount(accountRepository);

// interface adapter
new MainController(httpServer, signup, getAccount);

httpServer.listen(portToListen || portDefaultToListen);
