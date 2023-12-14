import dotenv from 'dotenv';
import pgp from 'pg-promise';
import type DatabaseConnection from './DatabaseConnection';

// Framework and Driver and Library
export default class PgPromiseAdapter implements DatabaseConnection {
  connection: any;

  constructor() {
    dotenv.config();
    const postgresDbUrl = process.env.POSTGRESDB_URL || '';
    this.connection = pgp()(postgresDbUrl);
  }

  query(statement: string, params: any): Promise<any> {
    return this.connection.query(statement, params);
  }

  async close(): Promise<void> {
    await this.connection.$pool.end();
  }
}
