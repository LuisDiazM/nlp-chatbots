import GoogleLoginButton from "../../auth/login";
import Navigation from "../../shared/navigation";
import logo from "../../assets/images/logo.png";

import { useUserLogin } from "../../auth/useUserLogin";
const HomePage = () => {
 const isUserLogin = useUserLogin() 

  return (
    <div>
      <Navigation></Navigation>
      <div style={{ display: "grid", justifyContent: "center" }}>
        <h1 style={{ fontSize: "42px" }}>Tech Talker</h1>
        <img
          src={logo}
          alt="logo"
          style={{ marginTop: "10px", marginBottom: "15px" }}
        ></img>
        {isUserLogin && <GoogleLoginButton></GoogleLoginButton>}
      </div>
    </div>
  );
};

export default HomePage;
