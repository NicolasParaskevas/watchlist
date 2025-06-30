export interface SymbolEntry {
  symbol: string;
}

export interface PriceUpdate {
  symbol: string;
  price: number;
}

export interface ClientMessage {
  action: 'subscribe' | 'unsubscribe';
  symbol: string;
}
