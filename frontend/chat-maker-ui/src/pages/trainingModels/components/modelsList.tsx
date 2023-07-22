import {Box} from "@mui/material";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import Divider from "@mui/material/Divider";
import InboxIcon from "@mui/icons-material/Inbox";
import TrainingIntentsModal from "./modals/trainingIntentsModal";
const ModelList = () => {

  return (
    <Box sx={{ width: "50%" }}>
        <List>
          <ListItem disablePadding>
            <ListItemButton>
              <ListItemIcon>
                <InboxIcon />
              </ListItemIcon>
              <ListItemText primary="Modelo 1" />
            </ListItemButton>
          </ListItem>
        </List>
      <Divider />
      <TrainingIntentsModal></TrainingIntentsModal>
    </Box>
  );
};

export default ModelList;
