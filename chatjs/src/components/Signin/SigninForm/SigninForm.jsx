import React, { useState } from "react";
import { useAuth } from "../../../hooks/useAuth";
import { Config } from "../../../config/config";
import "./SigninForm.scss"

const SigninForm = () => {
  const { login } = useAuth();
  const [formInfo, setFormInfo] = useState({
    userName: "",
    password: "",
  });
  const [errorText,setErrorText] = useState("");

  const handleInputChange = (event) => {
    setFormInfo({ ...formInfo, [event.target.name]:event.target.value });
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (formInfo.password === "") {
      setErrorText("パスワードを入力してください")
    }
    else if (formInfo.password.length < 6) {
      setErrorText("最低でもパスワードは6文字以上入力してください")
    }
    else {
      const requestOptions = {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({
          "name":formInfo.userName,
          "password": formInfo.password
        })
      };
      fetch(`${Config.postgresSource}/signin`, requestOptions)
      .then(resoponse => resoponse.json())
      .then(data => {
        if (Boolean(data["sign"])) {
          login(Boolean(data["sign"]), data["result"])
          setFormInfo({"userName": "", "password": ""});
          setErrorText("");
        } else {
          setErrorText(data["result"]);
        }
      });
    }
  }

  return (
    <div className="SigninForm">
      <form>
        <div className="InputField">
          <label>
            User Name
          </label>
          <input
            name="userName"
            type="text"
            value={formInfo.userName}
            onChange={handleInputChange}
          />
        </div>
        <div className="InputField">
          <label>
            Password
          </label>
          <input
            name="password"
            type="password"
            value={formInfo.password}
            onChange={handleInputChange}
          />
        </div>
        <p>{errorText}</p>
        <div className="action">
          <button type="button" onClick={handleSubmit}>Sign in</button>
        </div>
      </form>
      <div className="link"><a href="/">SignUp</a></div>
    </div>
  )
}

export default SigninForm