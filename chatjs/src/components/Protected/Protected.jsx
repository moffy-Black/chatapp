import { Navigate, useOutlet } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth";
import ProtectedHeader from "./ProtectedHeader";

const Protected = () => {
  const { sign } = useAuth();
  const outlet = useOutlet();

  if (!sign) {
    return <Navigate to="/" />;
  }

  return (
    <div>
      <ProtectedHeader />
      {outlet}
    </div>
  );
};

export default Protected;