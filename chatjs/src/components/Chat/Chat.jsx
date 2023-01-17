import React, { useState, useEffect, useRef } from "react";
import { useAuth } from "../../hooks/useAuth";
import { connect, sendMsg } from "./websocket";
import ChatHistory from "./ChatHistory";
import ChatInput from "./ChatInput";

const Chat = (props) => {
  const { userInfo } = useAuth();
  const scrollBottomRef = useRef(null);
  const [chatHistory, setChatHistory] = useState([])

  useEffect(() => {
    connect((msg) => {
      console.log("New Message")
      setChatHistory(prevState => (
        [...prevState, msg]
      ))
    });
  }, []);

  useEffect(() => {
    scrollBottomRef.current.scrollIntoView();
    console.log("scroll")
  }, [chatHistory])

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
    <div ref={scrollBottomRef}/>
    </div>
  )
}

export default Chat;
