import { useEffect, useState } from "react";
import { useSelector } from "react-redux";

export const useAuthCredential = (): string => {
  const authState = useSelector((store: any) => store.auth);
  const [token, setToken] = useState<string>("");

  useEffect(() => {
    const accessToken = authState.credential;
    setToken(accessToken);
  }, []);

  return token;
};
