import React from "react";

interface Props {
  symbols: string[];
  prices: Record<string, number>;
  onRemove: (symbol: string) => void;
}

const WatchlistTable: React.FC<Props> = ({ symbols, prices, onRemove }) => {
  return (
    <table>
      <thead>
        <tr>
          <th>Symbol</th>
          <th>Price</th>
          <th></th>
        </tr>
      </thead>
      <tbody>
        {symbols.map(symbol => (
          <tr key={symbol}>
            <td>{symbol}</td>
            <td>{prices[symbol] ?? "..."}</td>
            <td>
              <button onClick={() => onRemove(symbol)}>Remove</button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default WatchlistTable;
