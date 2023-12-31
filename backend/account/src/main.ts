import dotenv from 'dotenv';
import GetAccount from './application/usecase/GetAccount';
import Signup from './application/usecase/Signup';
import MainController from './infra/controller/MainController';
import PgPromiseAdapter from './infra/database/PgPromiseAdapter';
import Registry from './infra/di/Registry';
import ExpressAdapter from './infra/http/ExpressAdapter';
import LoggerConsole from './infra/logger/LoggerConsole';
import AccountRepositoryDatabase from './infra/repository/AccountRepositoryDatabase';

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

const registry = Registry.getInstance();
registry.register('httpServer', httpServer);
registry.register('signup', signup);
registry.register('getAccount', getAccount);

new MainController();
dotenv.config();
const PORT = parseInt(process.env.PORT || '3001');
httpServer.listen(PORT);
