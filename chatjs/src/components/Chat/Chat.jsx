import React, { useState, useEffect } from "react";
import { useAuth } from "../../hooks/useAuth";
import { connect, sendMsg } from "./websocket";
import ChatHistory from "./ChatHistory";
import ChatInput from "./ChatInput";

const Chat = (props) => {
  const { userInfo } = useAuth();
  const [chatHistory, setChatHistory] = useState([])

  useEffect(() => {
    connect((msg) => {
      console.log("New Message")
      setChatHistory(prevState => (
        [...prevState, msg]
      ))
    });
  }, []);

  const send = (text) => {
    const msg = {
      username: userInfo["name"],
      text: text
    }
    sendMsg(msg);
  }

  return (
    <div className="Chat">
    <ChatHistory handleAuthChange={props.handleAuthChange} chatHistory={chatHistory} />
    <ChatInput send={send}/>
    </div>
  )
}

export default Chat;
