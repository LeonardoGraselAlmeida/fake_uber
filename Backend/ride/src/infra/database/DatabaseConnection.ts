// Porta - Interface Adapters
export default interface DatabaseConnection {
  query(statement: string, params: unknown): Promise<any>;
  close(): Promise<void>;
}
