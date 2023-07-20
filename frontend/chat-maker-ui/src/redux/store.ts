import { configureStore } from "@reduxjs/toolkit";
import { userSlice } from "./states/user";

export const store = configureStore({
  reducer: {
    user: userSlice.reducer,
  },
});
