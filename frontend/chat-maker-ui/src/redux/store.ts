import { configureStore } from "@reduxjs/toolkit";
import { userSlice } from "./states/user";
import { authSlice } from "./states/auth";
import { messageSlice } from "./states/messages";

export const store = configureStore({
  reducer: {
    user: userSlice.reducer,
    auth: authSlice.reducer,
    messages: messageSlice.reducer,
  },
});
