import GoogleLoginButton from "../../auth/login";
import logo from "../../assets/images/logo.png";
import Navigation from "../../shared/navigation";

const LoginPage = () => {
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
        <GoogleLoginButton></GoogleLoginButton>
      </div>
    </div>
  );
};

export default LoginPage;
