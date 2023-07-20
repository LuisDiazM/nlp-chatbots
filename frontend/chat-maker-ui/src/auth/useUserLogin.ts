import { useEffect, useState } from "react";
import { useSelector } from "react-redux";

export const useUserLogin = (): boolean => {
  const userState = useSelector((store: any) => store.user);
  const [isUserLogin, setIsUserLogin] = useState<boolean>(false);

  useEffect(() => {
    if (userState?.name && userState?.email) {
      setIsUserLogin(false);
    } else {
      setIsUserLogin(true);
    }
  }, [userState]);

  return isUserLogin;
};
