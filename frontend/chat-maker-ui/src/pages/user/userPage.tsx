import { useSelector } from "react-redux";
import Navigation from "../../shared/navigation";
import UserChartsComponent from "./components/userChartsComponent";
import UserLicenseComponent from "./components/userLicenseComponent";
import styles from "./userPage.module.scss";
import { ACCESS_TOKEN } from "../../shared/utilities/constants";
import { LicensesModel } from "../../shared/utilities/models";
import useSWR from "swr";
import { fetcher } from "../../shared/utilities/fetcher";
import { environment } from "../../environments/environments";

const UserPage = () => {
    const userData = useSelector((store: any) => store.user);
    const token = sessionStorage.getItem(ACCESS_TOKEN);
    const userId = userData?.email ?? "";
  
    const { data } = useSWR<LicensesModel>(
      `${environment.BACKEND_URL}/license?userId=${userId}`,
      (url) => fetcher(url, token ?? "")
    );
  
  return (
    <div>
      <Navigation></Navigation>
      <div className={styles.userContainer}>
        <UserLicenseComponent license={data} user={userData} key={0}></UserLicenseComponent>
        <UserChartsComponent license={data} key={1}></UserChartsComponent>
      </div>
    </div>
  );
};

export default UserPage;
