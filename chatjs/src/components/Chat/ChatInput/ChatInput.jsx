import React from "react";
import { useEffect, useState } from "react";
import SpeechRecognition, { useSpeechRecognition } from 'react-speech-recognition';
import "./ChatInput.scss";
import MicIcon from "../../Icon/Mic";

const ChatInput = (props) => {
  const [chat, setChat] = useState("")
  const [isMic, setIsMic] = useState(false)
  const { interimTranscript, finalTranscript, resetTranscript } = useSpeechRecognition();

  useEffect(() => {
    if (interimTranscript !== '') {
      setChat(interimTranscript);
    }
  }, [interimTranscript]);

  useEffect(() => {
    if (finalTranscript !== '') {
      setChat(finalTranscript);
      props.send(finalTranscript);
    }
    resetTranscript();
  }, [finalTranscript]);

  const onMic = () => {
    setChat("");
    SpeechRecognition.startListening({ language: 'ja-JP', continuous:true });
    console.log("onMic");
    setIsMic(!isMic);
  }

  const offMic = () => {
    setChat("");
    SpeechRecognition.abortListening();
    resetTranscript();
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
              if (event.key === "Enter" && event.ctrlKey) {
                props.send(chat)
                event.target.value=""
              }
            }}
          />
        </div>
      }
    </div>
  )
}

export default ChatInput;