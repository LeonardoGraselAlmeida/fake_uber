import crypto from 'crypto';
import GetTransactionByRideId from '../src/application/usecase/GetTransactionByRideId';
import ProcessPayment from '../src/application/usecase/ProcessPayment';
import PgPromiseAdapter from '../src/infra/database/PgPromiseAdapter';
import Queue from '../src/infra/queue/Queue';
import TransactionRepositoryORM from '../src/infra/repository/TransactionRepositoryORM';

test('Deve processar um pagamento', async function () {
  const connection = new PgPromiseAdapter();
  const transactionRepository = new TransactionRepositoryORM(connection);
  const queue = new Queue();
  const processPayment = new ProcessPayment(transactionRepository, queue);
  const rideId = crypto.randomUUID();
  const inputProcessPayment = {
    rideId,
    creditCardToken: '123456789',
    amount: 1000
  };
  await processPayment.execute(inputProcessPayment);
  const getTransactionByRideId = new GetTransactionByRideId(
    transactionRepository
  );
  const output = await getTransactionByRideId.execute(rideId);
  expect(output.rideId).toBe(rideId);
  expect(output.status).toBe('paid');
  await connection.close();
});
