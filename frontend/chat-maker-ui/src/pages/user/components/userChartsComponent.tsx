import useSWR from "swr";
import { ACCESS_TOKEN } from "../../../shared/utilities/constants";
import { LicenseUsage, LicensesModel } from "../../../shared/utilities/models";
import styles from "./components.module.scss";
import { environment } from "../../../environments/environments";
import { fetcher } from "../../../shared/utilities/fetcher";
import { BarChart } from "@mui/x-charts/BarChart";
import {
  createDatasetsFromLicenseUsageRateLimits,
  createDatasetsFromLicenseUsageTrainigs,
} from "../utilities/createDatasetsFromLicenseUsage";
import Typography from "@mui/material/Typography";

interface UserChartsProps {
  license: LicensesModel | undefined;
}

const UserChartsComponent = ({ license }: UserChartsProps) => {
  const token = sessionStorage.getItem(ACCESS_TOKEN);

  const { data, isLoading } = useSWR<LicenseUsage>(
    `${environment.BACKEND_URL}/license/usage?licenseId=${license?.id}`,
    (url: string) => fetcher(url, token ?? "")
  );

  const datasetRateLimits = createDatasetsFromLicenseUsageRateLimits(data);
  const datasetTrainings = createDatasetsFromLicenseUsageTrainigs(data);
  return (
    <div className={styles.chartsContainer}>
      <Typography variant="h4">
        Consumo de la licencia por características
      </Typography>
      {!isLoading && datasetRateLimits?.length > 0 && (
        <BarChart
          dataset={datasetRateLimits}
          xAxis={[{ scaleType: "band", dataKey: "x" }]}
          series={[
            {
              label: "interacciones bot",
              dataKey: "y",
            },
          ]}
          width={500}
          height={300}
        />
      )}

      {!isLoading && datasetTrainings?.length>0 && (
        <BarChart
          dataset={datasetTrainings}
          xAxis={[
            {
              id: "barCategories",
              scaleType: "band",
              dataKey: "x",
            },
          ]}
          series={[
            {
              label: "Entrenamientos",
              color: "#29a543",
              dataKey: "y",
            },
          ]}
          width={500}
          height={300}
        />
      )}
    </div>
  );
};

export default UserChartsComponent;
