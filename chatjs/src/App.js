import React from "react";
import { Route, Routes } from "react-router-dom";
import "./App.css";
import Base from "./components/base";
import Chat from "./components/Chat/Chat";
import SignupForm from "./components/Signup/SignupForm";
import SigninForm from "./components/Signin/SigninForm";
import Protected from "./components/Protected";

function App() {
  return (
    <Routes>
      <Route element={<Base />}>
        <Route path={'/'} element={<SignupForm />} />
        <Route path={'/signin'} element={<SigninForm />} />
      </Route>

      <Route path={'/dashboard'} element={<Protected />} >
        <Route path={'chat'} element={<Chat />} />
      </Route >
    </Routes>
  )
}


export default App;
