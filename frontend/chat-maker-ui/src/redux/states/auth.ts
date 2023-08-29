import { createSlice } from "@reduxjs/toolkit";

export const initalAuthState = {
  credential: "",
  clientId: "",
};

export const authSlice = createSlice({
  name: "auth",
  initialState: { ...initalAuthState },
  reducers: {
    createAuth: (_, action) => {
      return action.payload;
    },
    resetAuth: ()=>{
        return {...initalAuthState}
    }
  },
});

export const {createAuth, resetAuth} = authSlice.actions