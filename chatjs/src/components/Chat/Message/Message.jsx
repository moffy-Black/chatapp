import React, { Component } from "react";
import "./Message.scss";

class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp
    };
  }

  render() {
    return (
      <div className="MessageArea">
        <p className="UserName">{this.state.message.body.username}</p>
        <div className="Message">{this.state.message.body.text}</div>
      </div>
    )
  }
}

export default Message;