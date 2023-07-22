import Navigation from "../../shared/navigation";
import logo from "../../assets/images/logo.png";

import { useUserLogin } from "../../auth/useUserLogin";
import { Button } from "@mui/material";
import { useNavigate } from "react-router-dom";
const HomePage = () => {
  const isUserLogin = useUserLogin();

  const navigate = useNavigate();

  const handleLogin = () => {
    navigate("/login");
  };

  const handleRegister = ()=>{
    navigate("/register")
  }

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

        {isUserLogin && (
          <>
            <Button variant="contained" color="success" onClick={handleLogin}>Iniciar sesión</Button>
            <Button onClick={handleRegister}> Registrarse</Button>
          </>
        )}
      </div>
    </div>
  );
};

export default HomePage;
