import { Config } from "../../../config/config";

var socket = new WebSocket(Config.websocketSource); // require minikube service port

let connect = cb => {
  console.log("connecting");

  socket.onopen = () => {
    console.log("Successfully Connected");
  };

  socket.onmessage = msg => {
    console.log(msg);
    cb(msg);
  };

  socket.onclose = msg => {
    console.log(msg);
    cb(msg);
  };

  socket.onerror = error => {
    console.log("Socket Error: ", error);
  };
};

let sendMsg = msg => {
  console.log("sending msg: ", msg);
  socket.send(JSON.stringify(msg));
};

export { connect, sendMsg };