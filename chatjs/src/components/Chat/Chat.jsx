import React, { useState, useEffect } from "react";
import { useAuth } from "../../hooks/useAuth";
import { connect, sendMsg } from "./websocket";
import ChatHistory from "./ChatHistory";
import ChatInput from "./ChatInput";

const Chat = (props) => {
  const { userInfo } = useAuth();
  const [chatHistory, setChatHIstory] = useState([])

  useEffect(() => {
    connect((msg) => {
      console.log("New Message")
      setChatHIstory(prevState => (
        [...prevState, msg]
      ))
    });
  }, []);

  const send = (event) => {
    if(event.keyCode === 13) {
      const msg = {
        username: userInfo["name"],
        text: event.target.value
      }
      sendMsg(msg);
      event.target.value = "";
    }
  }

  return (
    <div className="Chat">
    <ChatHistory handleAuthChange={props.handleAuthChange} chatHistory={chatHistory} />
    <ChatInput send={send}/>
    </div>
  )
}

export default Chat;
