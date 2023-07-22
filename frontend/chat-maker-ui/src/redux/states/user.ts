import { createSlice } from "@reduxjs/toolkit";

export const initalUserState = {
  name: "",
  email: "",
  picture: "",
  isLicenseValid: false,
};

export const userSlice = createSlice({
  name: "user",
  initialState: {
    ...initalUserState,
  },
  reducers: {
    createUser: (_, action) => {
      return action.payload;
    },
    updateUser: (state, action) => {
      return { ...state, ...action.payload };
    },
    resetUser: () => {
      return {
        ...initalUserState,
      };
    },
  },
});

export const { createUser, updateUser, resetUser } = userSlice.actions;
