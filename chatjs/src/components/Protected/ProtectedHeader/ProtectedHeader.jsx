import React from "react";
import { useAuth } from "../../../hooks/useAuth";
import "./ProtectedHeader.scss";

const ProtectedHeader = () => {
  const { logout } = useAuth();
  return(
    <div className="protectedHeader">
      <h2>WebSocket Chat App</h2>
      <div className="LogOut">
        <button type="button" onClick={() => {logout()}}>SIGN OUT</button>
      </div>
    </div>
  );
}

export default ProtectedHeader;