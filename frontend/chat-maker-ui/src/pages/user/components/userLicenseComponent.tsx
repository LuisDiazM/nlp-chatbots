import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import { LicensesModel } from "../../../shared/utilities/models";
import { Chip } from "@mui/material";
import DialogActions from "@mui/material/DialogActions";
import Dialog from "@mui/material/Dialog";
import DialogTitle from "@mui/material/DialogTitle";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import InboxIcon from "@mui/icons-material/Inbox";
import DraftsIcon from "@mui/icons-material/Drafts";
import { useState } from "react";
import { ACCESS_TOKEN } from "../../../shared/utilities/constants";
import { deleteUserAndTrainingModels } from "../../../shared/utilities/trainingModels";
import { useDispatch } from "react-redux";
import { logout } from "../../../auth/logout";
import { resetAuth } from "../../../redux/states/auth";
import { resetUser } from "../../../redux/states/user";
import { useNavigate } from "react-router-dom";

interface UserLicenseProps {
  user: any;
  license: LicensesModel | undefined;
}

const UserLicenseComponent = ({ user, license }: UserLicenseProps) => {
  const createLicense = new Date(license?.created_at ?? "");
  const expiredDate = new Date(license?.expired_at ?? "");
  const token = sessionStorage.getItem(ACCESS_TOKEN);
  const dispatcher = useDispatch();
  const navigate = useNavigate();

  const [openDeleteDialog, setOpenDeleteDialog] = useState(false);

  const handleCloseDeleteModel = () => {
    setOpenDeleteDialog(false);
  };

  const handleOpenDeleteModel = () => {
    setOpenDeleteDialog(true);
  };

  const handleDeleteModel = async () => {
    await deleteUserAndTrainingModels(token ?? "", user?.email ?? "");
    setOpenDeleteDialog(false);
      logout();
      dispatcher(resetAuth());
      dispatcher(resetUser());
      navigate("/");
    
  };
  return (
    <>
     <Card sx={{ maxWidth: 345 }}>
      <CardMedia
        component="img"
        alt="user"
        sx={{ width: 150, margin: "auto" }}
        image={user?.picture ?? ""}
      />
      <CardContent>
        <Typography gutterBottom variant="h5" component="div">
          {user?.name ?? ""}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Correo: {user?.email ?? ""}
        </Typography>
        <Typography variant="body2">
          <h3>Licencia</h3>
          <p>
            <b>Creado: </b>
            {`${createLicense.getFullYear()}-${createLicense.getMonth()}-${createLicense.getDate()}`}
          </p>
          <p>
            <b>Fecha de vencimiento: </b>
            {`${expiredDate.getFullYear()}-${
              expiredDate.getMonth() + 1
            }-${expiredDate.getDate()}`}
          </p>
          <Chip color="success" label={license?.type} />
        </Typography>
        <hr></hr>
        <Typography>Características disponibles</Typography>
        <List>
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <InboxIcon />
              </ListItemIcon>
              <ListItemText
                primary={`Entrenamientos: ${license?.features.trainings}`}
              />
            </ListItemButton>
          </ListItem>
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <DraftsIcon />
              </ListItemIcon>
              <ListItemText
                primary={`Interacciones bot: ${license?.features.rate_limit}`}
              />
            </ListItemButton>
          </ListItem>
        </List>
      </CardContent>
      <CardActions>
        <Button variant="contained" color="error" size="small" onClick={handleOpenDeleteModel}>
          Eliminar
        </Button>
      </CardActions>
    </Card>

    <Dialog
        open={openDeleteDialog}
        onClose={handleCloseDeleteModel}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">
          ¿Está seguro borrar el usuario?
          Esta acción no se puede deshacer, no se podrá registrar de nuevo de manera gratuita con este mismo usuario
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

export default UserLicenseComponent;
