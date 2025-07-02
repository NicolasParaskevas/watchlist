import React from "react";
import { SymbolEntry } from "../types";

interface Props {
  symbols: SymbolEntry[];
  prices: Record<string, number>;
  onRemove: (symbol: string) => void;
}

const WatchlistTable: React.FC<Props> = ({ symbols, prices, onRemove }) => {
  return (
    <table>
      <thead>
        <tr>
          <th>Symbol</th>
          <th>Name</th>
          <th>Price</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {symbols.map(symbol => (
          <tr key={symbol.symbol}>
            <td>{symbol.symbol}</td>
            <td>{symbol.name}</td>
            <td>{prices[symbol.symbol] ?? "..."}</td>
            <td>
              <button onClick={() => onRemove(symbol.symbol)}>Remove</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default WatchlistTable;
