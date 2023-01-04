import { createContext, useContext, useMemo } from "react";
import { useNavigate } from "react-router-dom";
import { useLocalStorage } from "./useLocalStorage";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [sign, setSign] = useLocalStorage("sign", false);
  const [userInfo, setUserInfo] = useLocalStorage("userInfo")
  const navigate = useNavigate();

  const login = async (sign, result) => {
    setSign(sign);
    setUserInfo(result);
    navigate("/dashboard/chat", { replace: true });
  };

  const logout = () => {
    setSign(false);
    setUserInfo(null);
    navigate("/", { replace: true });
  };

  const value = useMemo(
    () => ({
      sign,
      userInfo,
      login,
      logout
    }),
    [sign]
  );

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuth = () => {
  return useContext(AuthContext);
};
