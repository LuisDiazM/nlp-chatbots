import { configureStore } from "@reduxjs/toolkit";
import { userSlice } from "./states/user";
import { authSlice } from "./states/auth";

export const store = configureStore({
  reducer: {
    user: userSlice.reducer,
    auth: authSlice.reducer,
  },
});
