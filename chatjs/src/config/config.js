const Config = {
  websocketSource: process.env.REACT_APP_WEBSOCKET_SOURCE ?? "http://localhost:8080",
  postgresSource: process.env.REACT_APP_POSTGRES_SOURCE ?? "http://localhost:8000"
};

export {Config};