import {
  Button,
  Card,
  CardActions,
  CardContent,
  Typography,
} from "@mui/material";
import {
  TrainingModel,
  deleteTrainingModel,
} from "../../utilities/trainingModels";
import ModelData from "./modelData";
import { useState } from "react";

import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogTitle from "@mui/material/DialogTitle";
import { ACCESS_TOKEN } from "../../../shared/utilities/constants";

interface TrainingModelProps {
  model: TrainingModel;
}

const TrainingModelCard: React.FC<TrainingModelProps> = ({ model }) => {
  const [modalOpen, setModalOpen] = useState(false);
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);


  const token = sessionStorage.getItem(ACCESS_TOKEN);

  const handleOpenModal = () => {
    setModalOpen(true);
  };

  const handleCloseModal = (): void => {
    setModalOpen(false);
  };

  const handleDeleteModel = async () => {
    await deleteTrainingModel(token ?? "", model?.id ?? "");
    setOpenDeleteDialog(false);
  };

  const handleCloseDeleteModel = () => {
    setOpenDeleteDialog(false);
  };

  const handleOpenDeleteModel = () => {
    setOpenDeleteDialog(true);
  };

  return (
    <>
      <Card sx={{ minWidth: 275, marginLeft: "10px", marginBottom: "10px" }}>
        <CardContent>
          <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
            {model?.title}
          </Typography>
          <Typography variant="body2">
            {model?.description}
            <br />
          </Typography>
        </CardContent>
        <CardActions>
          <Button size="small" onClick={handleOpenModal}>
            Editar
          </Button>
          <Button size="small" color="error" onClick={handleOpenDeleteModel}>
            Borrar
          </Button>
        </CardActions>
      </Card>
      <ModelData
        isOpen={modalOpen}
        onClose={setModalOpen}
        model={model}
        handleClickClose={handleCloseModal}
      ></ModelData>

      <Dialog
        open={openDeleteDialog}
        onClose={handleCloseDeleteModel}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">
          ¿Está seguro borrar el modelo?
        </DialogTitle>
        <DialogActions>
          <Button onClick={handleCloseDeleteModel}>No borrar</Button>
          <Button color="error" onClick={handleDeleteModel} autoFocus>
            Borrar
          </Button>
        </DialogActions>
      </Dialog>
    </>
  );
};

export default TrainingModelCard;
