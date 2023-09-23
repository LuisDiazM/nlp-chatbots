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
} from "../../../shared/utilities/trainingModels";
import ModelData from "./modelData";
import { useState } from "react";

import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogTitle from "@mui/material/DialogTitle";
import { ACCESS_TOKEN } from "../../../shared/utilities/constants";
import { environment } from "../../../environments/environments";
import useSWR from "swr";
import { fetcher } from "../../../shared/utilities/fetcher";
import TestingBotModal from "./modals/testingChatbot/testingBotContainer";

interface TrainingModelProps {
  model: TrainingModel;
}

interface NNModels {
  created: Date;
  model_name: string;
  nlu_intent_id: string;
  user_id: string;
  id: string;
  bucket_name: string;
}

const TrainingModelCard: React.FC<TrainingModelProps> = ({ model }) => {
  const [modalOpen, setModalOpen] = useState(false);
  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
  const [openTestingBot, setOpenTestingBot] = useState(false);
const [modelId, setModelId] = useState<string>("")

  const handleClose = () => setOpenTestingBot(false);

  const token = sessionStorage.getItem(ACCESS_TOKEN);

  const { data, isLoading } = useSWR<NNModels[]>(
    `${environment.BACKEND_URL}/nn-models?trainingId=${model?.id}`,
    (url) => fetcher(url, token ?? "")
  );

  const handleTestBot = () => {
    const modelsSortedDescending = data?.sort(
      (a, b) => Number(a.created) - Number(b.created)
    );
    if (modelsSortedDescending) {
      const lastModel = modelsSortedDescending[modelsSortedDescending?.length - 1]
      setOpenTestingBot(true);
      setModelId(lastModel.id)
    }
  };

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
          {!isLoading && (
            <Button size="small" color="secondary" onClick={handleTestBot}>
              {" "}
              Probar bot
            </Button>
          )}
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

      <TestingBotModal
        closeTesting={handleClose}
        isOpenTesting={openTestingBot}
        modelId={modelId}
      ></TestingBotModal>
    </>
  );
};

export default TrainingModelCard;
