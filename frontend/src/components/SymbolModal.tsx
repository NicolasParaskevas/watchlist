import React, { useEffect, useState } from "react";
import axios from "axios";
import { SymbolEntry } from "../types";

interface Props {
  onClose: () => void;
  onAdd: (symbol: string) => void;
}

const SymbolModal: React.FC<Props> = ({ onClose, onAdd }) => {
  const [symbols, setSymbols] = useState<SymbolEntry[]>([]);

  useEffect(() => {
    axios.get("http://localhost:8080/symbols-list").then(res => {
      setSymbols(res.data);
    });
  }, []);

  return (
    <div className="modal">
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
  );
};

export default SymbolModal;
