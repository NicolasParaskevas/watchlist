export interface SymbolEntry {
  symbol: string;
  name: string;
}

export interface PriceUpdate {
  symbol: string;
  price: number;
}

export interface ClientMessage {
  action: 'subscribe' | 'unsubscribe';
  symbol: string;
}
