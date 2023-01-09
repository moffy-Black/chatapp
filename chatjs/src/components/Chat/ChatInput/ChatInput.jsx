import React from "react";
import { useState } from "react";
import "./ChatInput.scss";
import MicIcon from "../../Icon/Mic";

const ChatInput = (props) => {
  const [chat, setChat] = useState("")
  const [isMic, setIsMic] = useState(false)

  const onMic = () => {
    setChat("");
    console.log("onMic");
    setIsMic(!isMic);
  }

  const offMic = () => {
    setChat("");
    console.log("offMic");
    setIsMic(!isMic);
  }

  return (
    <div className="ChatInput">
      {isMic ?
        <div className="chatKapsel">
          <button onClick={offMic}><MicIcon fill="#01579B" width={20} height={30}/></button>
          <p className="chat">{chat}</p>
        </div>:
        <div className="chatKapsel">
          <button onClick={onMic}><MicIcon fill="#ffffff" width={20} height={30}></MicIcon></button>
          <input className="chat"
            onChange={
              (event) => {
                setChat(event.target.value);
              }
            }
            onKeyDown={(event) => {
              if ((event.ctrlKey && !event.metaKey) || (!event.ctrlKey && event.metaKey) || (event.key === "Enter")) {
                props.send(chat)
              }
            }}
          />
        </div>
      }
    </div>
  )
}

export default ChatInput;