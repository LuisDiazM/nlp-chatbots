import { GoogleLogin } from "@react-oauth/google";
import { useDispatch } from "react-redux";
import { createUser } from "../redux/states/user";
import { useNavigate } from "react-router-dom";

const GoogleLoginButton = () => {
  const dispatcher = useDispatch();
  const navigate = useNavigate();
  const handleSuccess = (response: any) => {
    dispatcher(
      createUser({ name: response.credential, email: response.clientId })
    );
    navigate("/models");
  };

  const handleFailure = () => {
    console.error("Login Failure");
  };

  return <GoogleLogin onSuccess={handleSuccess} onError={handleFailure} />;
};

export default GoogleLoginButton;
