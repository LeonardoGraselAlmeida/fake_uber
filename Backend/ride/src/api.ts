import type { Request, Response } from 'express';
import express from 'express';
import AccountDAODatabase from './AccountDAODatabase';
import GetAccount from './GetAccount';
import LoggerConsole from './LoggerConsole';
import Signup from './Signup';

const app = express();
app.use(express.json());

app.post('/signup', async function (req: Request, res: Response) {
  try {
    const input = req.body;
    const accountDAO = new AccountDAODatabase();
    const logger = new LoggerConsole();
    const signup = new Signup(accountDAO, logger);
    const output = await signup.execute(input);
    res.json(output);
  } catch (e: any) {
    res.status(422).json({
      message: e.message
    });
  }
});

app.get('/accounts/:accountId', async function (req: Request, res: Response) {
  const accountId = req.params.accountId;
  const accountDAO = new AccountDAODatabase();
  const getAccount = new GetAccount(accountDAO);
  const output = await getAccount.execute(accountId);
  res.json(output);
});

app.listen(3000);
