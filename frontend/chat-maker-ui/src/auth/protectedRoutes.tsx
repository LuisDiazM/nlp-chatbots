import { Outlet } from "react-router-dom";
import { useUserLogin } from "./useUserLogin";
import HomePage from "../pages/home/homePage";

const ProtectedRoutes = () => {
  const isUserLogin = useUserLogin();
  return !isUserLogin ? <Outlet /> : <HomePage />;
};

export default ProtectedRoutes;
