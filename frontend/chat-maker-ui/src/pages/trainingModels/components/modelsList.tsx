import { Box, CircularProgress } from "@mui/material";

import TrainingIntentsModal from "./modals/trainingIntentsModal";
import { useSelector } from "react-redux";
import { TrainingModel } from "../../utilities/trainingModels";
import TrainingModelCard from "./modelCard";
import useSWR from "swr";
import { fetcher } from "../../utilities/fetcher";
import { environment } from "../../../environments/environments";
import { ACCESS_TOKEN } from "../../../shared/utilities/constants";

const ModelList = () => {
  const userState = useSelector((store: any) => store.user);
  const token = sessionStorage.getItem(ACCESS_TOKEN);
  const userId = userState.email;
  const { data, isLoading } = useSWR<TrainingModel[]>(
    `${environment.BACKEND_URL}/training-model?userId=${userId}`,
    (url) => fetcher(url, token ?? "")
  );

  return (
    <Box sx={{ width: "50%" }}>
      <h2>Mis modelos</h2>
      <TrainingIntentsModal></TrainingIntentsModal>

      {isLoading && <CircularProgress color="inherit" />}
      {!isLoading && (
        <div>
          {data?.map((model, index) => {
            return (
              <TrainingModelCard
                key={model?.id ?? index}
                model={model}
              ></TrainingModelCard>
            );
          })}
        </div>
      )}
    </Box>
  );
};

export default ModelList;
