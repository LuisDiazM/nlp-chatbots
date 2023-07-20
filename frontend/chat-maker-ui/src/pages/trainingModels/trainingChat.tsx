import ModelList from "./components/modelsList";
import Navigation from "../../shared/navigation";

const TrainingChatPage = () => {
  return (
    <div>
      <Navigation></Navigation>
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <ModelList></ModelList>
      </div>
    </div>
  );
};

export default TrainingChatPage;
