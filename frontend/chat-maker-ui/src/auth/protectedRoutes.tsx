import { Outlet } from "react-router-dom";
import { useUserLogin } from "./useUserLogin";
import HomePage from "../pages/home/homePage";
import swal from "sweetalert";

const ProtectedRoutes = () => {
  const isUserLogin = useUserLogin();
  if (isUserLogin) {
    swal(
      "No tienes acceso!",
      "Tu usuario no está registrado o la licencia de uso está vencida!",
      "error"
    );
  }
  return !isUserLogin ? <Outlet /> : <HomePage />;
};

export default ProtectedRoutes;
