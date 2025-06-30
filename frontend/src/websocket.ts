import { ClientMessage, PriceUpdate } from "./types";

let socket: WebSocket | null = null;

export function connectWS(onPriceUpdate: (update: PriceUpdate) => void) {
  if (socket) return; // skip if we already have a connection
  socket = new WebSocket("ws://localhost:8080/ws");

  socket.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    if (msg.symbol && typeof msg.price === 'number') {
      onPriceUpdate(msg);
    }
  };

  socket.onclose = () => {
    console.log("[WebSocket] Disconnected");
    socket = null;
  };
}
export function sendMessage(msg: ClientMessage) {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(msg));
  }
}
