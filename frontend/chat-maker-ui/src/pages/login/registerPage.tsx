import { useDispatch } from "react-redux";
import logo from "../../assets/images/logo.png";
import { getUserLogin, registerUser } from "../../auth/utilities/fetchUser";
import Navigation from "../../shared/navigation";
import { GoogleLogin } from "@react-oauth/google";
import { useNavigate } from "react-router-dom";
import { createUser } from "../../redux/states/user";
import swal from "sweetalert";

const RegisterPage = () => {
  const dispatcher = useDispatch();
  const navigate = useNavigate();

  const handleFailure = () => {
    console.error("Login Failure");
  };

  const handleRegistry = (response: any) => {
    registerUser(response.credential).then((userCreated) => {
        if(!userCreated.is_created){
            swal(
                `Usuario`,
                "ya se encuentra registrado",
                "warning"
              );
        }
      getUserLogin(response.credential).then(({ user, is_licence_valid }) => {
        dispatcher(
          createUser({
            name: user.name,
            email: user.email,
            picture: user.picture,
            isLicenseValid: is_licence_valid,
          })
        );
        navigate("/models");
      });
    });
  };
  return (
    <div>
      <Navigation></Navigation>
      <div style={{ display: "grid", justifyContent: "center" }}>
        <h1 style={{ fontSize: "42px" }}>Tech Talker</h1>
        <img
          src={logo}
          alt="logo"
          style={{ marginTop: "10px", marginBottom: "15px" }}
        ></img>{" "}
        <h2>Registrarse con:</h2>
        <GoogleLogin onSuccess={handleRegistry} onError={handleFailure} />
      </div>
    </div>
  );
};

export default RegisterPage;
