import dotenv from 'dotenv';
import RequestRide from './application/usecase/RequestRide';
import SendReceipt from './application/usecase/SendReceipt';
import MainController from './infra/controller/MainController';
import PgPromiseAdapter from './infra/database/PgPromiseAdapter';
import Registry from './infra/di/Registry';
import AccountGatewayHttp from './infra/gateway/AccountGatewayHttp';
import ExpressAdapter from './infra/http/ExpressAdapter';
import LoggerConsole from './infra/logger/LoggerConsole';
import Queue from './infra/queue/Queue';
import QueueController from './infra/queue/QueueController';
import RideRepositoryDatabase from './infra/repository/RideRepositoryDatabase';

// composition root ou entry point
// criar o grafo de dependências utilizado no projeto

// framework and driver and library
const httpServer = new ExpressAdapter();
const databaseConnection = new PgPromiseAdapter();
const queue = new Queue();
const rideRepository = new RideRepositoryDatabase(databaseConnection);
const accountGateway = new AccountGatewayHttp();

// interface adapter
const logger = new LoggerConsole();

// use case
const sendReceipt = new SendReceipt();
const requestRide = new RequestRide(rideRepository, accountGateway, logger);

const registry = Registry.getInstance();
registry.register('httpServer', httpServer);
registry.register('queue', queue);
registry.register('sendReceipt', sendReceipt);
registry.register('requestRide', requestRide);

new MainController();
new QueueController();

dotenv.config();
const PORT = parseInt(process.env.PORT2 || '3000');
httpServer.listen(PORT);
