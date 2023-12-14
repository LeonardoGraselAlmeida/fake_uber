import dotenv from 'dotenv';
import GetTransactionByRideId from './application/usecase/GetTransactionByRideId';
import ProcessPayment from './application/usecase/ProcessPayment';
import MainController from './infra/controller/MainController';
import PgPromiseAdapter from './infra/database/PgPromiseAdapter';
import Registry from './infra/di/Registry';
import ExpressAdapter from './infra/http/ExpressAdapter';
import Queue from './infra/queue/Queue';
import QueueController from './infra/queue/QueueController';
import TransactionRepositoryORM from './infra/repository/TransactionRepositoryORM';

// composition root ou entry point
// criar o grafo de dependências utilizado no projeto

// framework and driver and library
const httpServer = new ExpressAdapter();
const databaseConnection = new PgPromiseAdapter();
const transactionRepository = new TransactionRepositoryORM(databaseConnection);

const queue = new Queue();
// interface adapter
// const logger = new LoggerConsole();

// use case
const processPayment = new ProcessPayment(transactionRepository, queue);
const getTransactionByRideId = new GetTransactionByRideId(
  transactionRepository
);

const registry = Registry.getInstance();
registry.register('httpServer', httpServer);
registry.register('queue', queue);
registry.register('processPayment', processPayment);
registry.register('getTransactionByRideId', getTransactionByRideId);

new MainController();
new QueueController();

dotenv.config();
const PORT = parseInt(process.env.PORT || '3002');
httpServer.listen(PORT);
