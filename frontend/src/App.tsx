import { useEffect, useState } from "react";
import axios from "axios";
import { connectWS, sendMessage } from "./websocket";
import { PriceUpdate, SymbolEntry } from "./types";
import SymbolModal from "./components/SymbolModal";
import WatchlistTable from "./components/WatchlistTable";

function App() {
  const [watchlist, setWatchlist] = useState<string[]>([]);
  const [prices, setPrices] = useState<Record<string, number>>({});
  const [showModal, setShowModal] = useState(false);
  const [symbols, setSymbols] = useState<SymbolEntry[]>([]);

  useEffect(() => {
    connectWS((update: PriceUpdate) => {
      setPrices(prev => ({ ...prev, [update.symbol]: update.price }));
    });
  }, []);

  useEffect(() => {
    axios.get("http://localhost:8080/symbols-list").then(res => {
      setSymbols(res.data);
    });
  }, []);

  const subscribe = (symbol: string) => {
    sendMessage({ action: "subscribe", symbol });
    setWatchlist(prev => [...new Set([...prev, symbol])]);
  };

  const unsubscribe = (symbol: string) => {
    sendMessage({ action: "unsubscribe", symbol });
    setWatchlist(prev => prev.filter(s => s !== symbol));
    setPrices(prev => {
      const updated = { ...prev };
      delete updated[symbol];
      return updated;
    });
  };

  return (
    <div>
      <h1>Watchlist</h1>
      <button onClick={() => setShowModal(true)}>Add Symbol</button>
      <WatchlistTable
        symbols={watchlist}
        prices={prices}
        onRemove={unsubscribe}
      />
      {showModal && (
        <SymbolModal
          onClose={() => setShowModal(false)}
          onAdd={subscribe}
          symbols={symbols}
        />
      )}
    </div>
  );
}

export default App;
