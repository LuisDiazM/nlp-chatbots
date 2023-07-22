import { environment } from "../../environments/environments";
import { UserCreated, UserWithLicenseValidation } from "../models/user.model";

export const getUserLogin = async (
  idToken: string
): Promise<UserWithLicenseValidation> => {
  let user: UserWithLicenseValidation = {
    is_licence_valid: false,
    user: { created_at: new Date(), email: "", id: "", name: "", picture: "" },
  };
  const url = new URL(`${environment.BACKEND_URL}/login`);
  const params = { id: idToken };
  url.search = new URLSearchParams(params).toString();
  const response = await fetch(url);
  if (response.status === 200) {
    user = await response.json();
  }
  return user;
};

export const registerUser = async (idToken: string): Promise<UserCreated> => {
  let userCreated: UserCreated = { is_created: false, user_id: null };
  const url = new URL(`${environment.BACKEND_URL}/register`);
  const params = { id: idToken };
  url.search = new URLSearchParams(params).toString();
  const response = await fetch(url, { method: "POST" });
  if (response.status === 200 || response.status === 201) {
    userCreated = await response.json();
  }
  return userCreated;
};
