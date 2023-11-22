import type { IDatabase } from 'pg-promise';
import pgp from 'pg-promise';
import type DatabaseConnection from './DatabaseConnection';

// Framework and Driver and Library
export default class PgPromiseAdapter implements DatabaseConnection {
  connection: IDatabase<undefined>;

  constructor() {
    this.connection = pgp()('postgres://postgres:123456@localhost:5432/app');
  }

  query(statement: string, params: unknown): Promise<unknown> {
    return this.connection.query(statement, params);
  }

  async close(): Promise<void> {
    await this.connection.$pool.end();
  }
}
