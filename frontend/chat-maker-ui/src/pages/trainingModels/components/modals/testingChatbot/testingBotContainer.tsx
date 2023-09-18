import { Box, Modal } from "@mui/material";
import TestingBotForm from "./testingBotForm";
import MessagesBlock from "./messages/messagesBlock";
import { useDispatch } from "react-redux";
import { resetMessages } from "../../../../../redux/states/messages";

const style = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  border: "2px solid #000",
  boxShadow: 24,
  p: 4,
};

interface TestingBotProps {
  isOpenTesting: boolean;
  closeTesting: () => void;
  modelId: string;
}

const TestingBotModal = ({
  isOpenTesting,
  closeTesting,
  modelId,
}: TestingBotProps) => {


  const dispatcher = useDispatch()

  const handleCloseTestingBot = ()=>{
    closeTesting()
    dispatcher(resetMessages())
  }

  return (
    <Modal
      open={isOpenTesting}
      onClose={handleCloseTestingBot}
      aria-labelledby="modal-modal-title"
      aria-describedby="modal-modal-description"
    >
      <Box sx={style}>
        <MessagesBlock></MessagesBlock>
        <TestingBotForm modelId={modelId}></TestingBotForm>
      </Box>
    </Modal>
  );
};

export default TestingBotModal;
