# watchlist
Stock terminal

## Tech Stack:
The project was built in Go v1.22.2 and React with Typescript. For the database it loads the symbols from a simple json file  
I did this for the sake of simplicity but I utilized the Repository pattern and have a dummy postgres reposiory that could be connected
to a real database

I don't have specific endpoints for subscribe/unsubscribe, the logic is handled through the websocket connection

## Instalation Guide
For the backend:
```
cd backend
make
```

For the frontend:
```
cd frontend
npm install
npm start
```