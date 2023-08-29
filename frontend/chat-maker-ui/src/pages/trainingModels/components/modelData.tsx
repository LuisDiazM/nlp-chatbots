import { Box, Modal } from "@mui/material";
import React from "react";
import { TrainingModel } from "../../utilities/trainingModels";
import TrainingChatbotForm from "./forms/trainingForm";

const style = {
  position: "absolute" as "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: "80%",
  bgcolor: "background.paper",
  border: "2px solid #000",
  boxShadow: 24,
  p: 4,
};

interface OpenModal {
  isOpen: boolean;
  onClose: React.Dispatch<React.SetStateAction<boolean>>;
  model: TrainingModel;
  handleClickClose: () => void;
}

const ModelData: React.FC<OpenModal> = ({
  isOpen,
  onClose,
  model,
  handleClickClose,
}) => {
  return (
    <div>
      <Modal
        open={isOpen}
        onClose={handleClickClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <TrainingChatbotForm
            model={model}
            isEdit={true}
            closeModal={onClose}
          ></TrainingChatbotForm>
        </Box>
      </Modal>
    </div>
  );
};

export default ModelData;
