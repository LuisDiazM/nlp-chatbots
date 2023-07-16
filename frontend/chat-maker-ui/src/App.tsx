import "./App.css";
import TrainingChatPage from "./pages/trainingChat";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
function App() {
  const router = createBrowserRouter([
    {
      path: "/train-chatbot",
      element: <TrainingChatPage></TrainingChatPage>,
    },
  ]);

  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
