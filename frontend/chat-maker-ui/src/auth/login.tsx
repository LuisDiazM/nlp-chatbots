import { GoogleLogin } from "@react-oauth/google";
import { useDispatch } from "react-redux";
import { createUser } from "../redux/states/user";
import { useNavigate } from "react-router-dom";
import { getUserLogin } from "./utilities/fetchUser";

const GoogleLoginButton = () => {
  const dispatcher = useDispatch();
  const navigate = useNavigate();
  const handleSuccess = async (response: any) => {
    getUserLogin(response.credential)
      .then(({ user, is_licence_valid }) => {
        dispatcher(
          createUser({
            name: user.name,
            email: user.email,
            picture: user.picture,
            isLicenseValid: is_licence_valid,
          })
        );
        navigate("/models");
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const handleFailure = () => {
    console.error("Login Failure");
  };

  return (
    <GoogleLogin
      onSuccess={handleSuccess}
      onError={handleFailure}
    />
  );
};

export default GoogleLoginButton;
