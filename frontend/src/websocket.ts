import { ClientMessage, PriceUpdate } from "./types";

let socket: WebSocket;

export function connectWS(onPriceUpdate: (update: PriceUpdate) => void) {
  socket = new WebSocket("ws://localhost:8080/ws");

  socket.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    if (msg.symbol && msg.price) {
      onPriceUpdate(msg);
    }
  };
}

export function sendMessage(msg: ClientMessage) {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(msg));
  }
}
