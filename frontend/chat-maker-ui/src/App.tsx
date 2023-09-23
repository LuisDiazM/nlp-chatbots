import "./App.css";
import {
  BrowserRouter as Router,
  Route,
  Routes,
} from "react-router-dom";
import { Provider } from "react-redux";
import { store } from "./redux/store";
import { GoogleOAuthProvider } from "@react-oauth/google";
import { environment } from "./environments/environments";
import HomePage from "./pages/home/homePage";
import ProtectedRoutes from "./auth/protectedRoutes";
import TrainingChatPage from "./pages/trainingModels/trainingChat";
import LoginPage from "./pages/login/loginPage";
import RegisterPage from "./pages/login/registerPage";
import UserPage from "./pages/user/userPage";

function App() {
  return (
    <>
      <Provider store={store}>
        <GoogleOAuthProvider clientId={environment.CLIENT_ID}>
          <Router>
            <Routes>
              <Route path="/" element={<HomePage></HomePage>} />
              <Route path="/login" element={<LoginPage></LoginPage>} />
              <Route
                path="/register"
                element={<RegisterPage></RegisterPage>}
              ></Route>
              <Route
                path="/models"
                element={<TrainingChatPage></TrainingChatPage>}
              ></Route>
              <Route path="/user" element={<UserPage></UserPage>}>
                {" "}
              </Route>
              <Route path="*" element={<HomePage></HomePage>} />{" "}
            </Routes>
          </Router>
        </GoogleOAuthProvider>
      </Provider>
    </>
  );
}

export default App;
