import React from "react";
import { SymbolEntry } from "../types";

interface Props {
  onClose: () => void;
  onAdd: (symbol: string) => void;
  symbols: SymbolEntry[];
}

const SymbolModal: React.FC<Props> = ({ onClose, onAdd, symbols }) => {
  return (
    <div className="modal">
      <div>
        <h2>Select a symbol</h2>
        <ul>
          {symbols.map(s => (
            <li key={s.symbol}>
              {s.symbol}
              <button onClick={() => { onAdd(s.symbol); onClose(); }}>Add</button>
            </li>
          ))}
        </ul>
        <button onClick={onClose}>Close</button>
      </div>
    </div>
  );
};

export default SymbolModal;
