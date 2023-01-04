import React from "react";
import { useOutlet } from "react-router-dom";
import Header from "./Header";

const Base = () => {
  const outlet = useOutlet();
  return (
    <div className="Base">
      <Header />
      {outlet}
    </div>
  )
}

export default Base;