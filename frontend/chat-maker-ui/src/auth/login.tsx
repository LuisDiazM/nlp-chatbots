import { GoogleLogin } from "@react-oauth/google";
import { useDispatch } from "react-redux";
import { createUser } from "../redux/states/user";
import { useNavigate } from "react-router-dom";
import { getUserLogin } from "./utilities/fetchUser";
import { createAuth } from "../redux/states/auth";
import { ACCESS_TOKEN } from "../shared/utilities/constants";

const GoogleLoginButton = () => {
  const dispatcher = useDispatch();
  const navigate = useNavigate();
  const handleSuccess = async (response: any) => {
    dispatcher(
      createAuth({
        credential: response.credential,
        clientId: response.clientId,
      })
    );
    sessionStorage.setItem(ACCESS_TOKEN, response.credential)
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

  return <GoogleLogin onSuccess={handleSuccess} onError={handleFailure} />;
};

export default GoogleLoginButton;
