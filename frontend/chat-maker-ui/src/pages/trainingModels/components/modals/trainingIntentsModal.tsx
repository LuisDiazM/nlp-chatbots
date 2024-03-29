import {Box} from "@mui/material";
import Button from "@mui/material/Button";
import Modal from "@mui/material/Modal";
import DialogContent from "@mui/material/DialogContent";
import { useState } from "react";
import React from "react";
import TrainingChatbotForm from "../forms/trainingForm";
import { TrainingModel } from "../../../../shared/utilities/trainingModels";
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



const TrainingIntentsModal: React.FC = () => {
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const initialValues: TrainingModel = {
    user_id: "",
    intents: [{ patterns: ["", "", "", "", ""], responses: [""], tag: "" }],
    title: "",
    description:""
  };
  return (
    <div>
      <Button style={{marginBottom:"10px"}} onClick={handleOpen} variant="contained" color="success">
        Crear Modelo
      </Button>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <DialogContent>
            <div style={{ maxHeight: "700px", overflowY: "auto", paddingTop:"10px" }}>
              <TrainingChatbotForm model={initialValues} isEdit={false} key={0} closeModal={setOpen}></TrainingChatbotForm>
            </div>
          </DialogContent>
        </Box>
      </Modal>
    </div>
  );
};

export default TrainingIntentsModal;
