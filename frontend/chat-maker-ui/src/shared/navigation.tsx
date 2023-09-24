import { Box } from "@mui/material";
import BottomNavigation from "@mui/material/BottomNavigation";
import BottomNavigationAction from "@mui/material/BottomNavigationAction";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import TopicIcon from "@mui/icons-material/Topic";
import LogoutIcon from "@mui/icons-material/Logout";
import { Link, useNavigate } from "react-router-dom";
import HomeIcon from "@mui/icons-material/Home";
import { useState } from "react";
import { useUserLogin } from "../auth/useUserLogin";
import { logout } from "../auth/logout";
import { useDispatch } from "react-redux";
import { resetUser } from "../redux/states/user";
import { resetAuth } from "../redux/states/auth";

const Navigation = () => {
  const [value, setValue] = useState(0);
  const navigate = useNavigate();
  const isUserLogin = useUserLogin();
  const dispatcher = useDispatch();
  const handleLogoutApp = () => {
    logout();
    dispatcher(resetAuth());
    dispatcher(resetUser());
    navigate("/");
  };

  return (
    <div>
      <Box sx={{ width: 500, margin: "auto" }}>
        <BottomNavigation
          showLabels
          value={value}
          onChange={(_, newValue) => {
            setValue(newValue);
          }}
        >
          <BottomNavigationAction
            icon={
              <Link to="/">
                <HomeIcon />
              </Link>
            }
          />
          {!isUserLogin && (
            <BottomNavigationAction
              label="Mis modelos"
              icon={
                <Link to="/models">
                  {" "}
                  <TopicIcon />
                </Link>
              }
            />
          )}
          {!isUserLogin && (
            <BottomNavigationAction
              label="Usuario"
              icon={
                <Link to={"/user"}>
                  <AccountCircleIcon />
                </Link>
              }
            />
          )}

          {!isUserLogin && (
            <BottomNavigationAction
              onClick={handleLogoutApp}
              label="Salir"
              icon={<LogoutIcon />}
            />
          )}
        </BottomNavigation>
      </Box>
    </div>
  );
};

export default Navigation;
