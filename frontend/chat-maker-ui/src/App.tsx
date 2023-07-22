import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { Provider } from "react-redux";
import { store } from "./redux/store";
import { GoogleOAuthProvider } from "@react-oauth/google";
import { environment } from "./environments/environments";
import HomePage from "./pages/home/homePage";
import ProtectedRoutes from "./auth/protectedRoutes";
import TrainingChatPage from "./pages/trainingModels/trainingChat";
import LoginPage from "./pages/login/loginPage";
import RegisterPage from "./pages/login/registerPage";

function App() {
  return (
    <>
      <Provider store={store}>
        <GoogleOAuthProvider clientId={environment.CLIENT_ID}>
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<HomePage></HomePage>} />
              <Route path="/login" element={<LoginPage></LoginPage>} />
              <Route
                path="/register"
                element={<RegisterPage></RegisterPage>}
              ></Route>
              <Route path="*" element={<HomePage></HomePage>} />{" "}
              <Route element={<ProtectedRoutes></ProtectedRoutes>}>
                <Route
                  path="/models"
                  element={<TrainingChatPage></TrainingChatPage>}
                ></Route>
              </Route>
            </Routes>
          </BrowserRouter>
        </GoogleOAuthProvider>
      </Provider>
    </>
  );
}

export default App;
