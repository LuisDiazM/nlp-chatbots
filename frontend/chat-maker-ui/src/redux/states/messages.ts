import { createSlice } from "@reduxjs/toolkit";

export enum OriginMessage {
  BOT = "BOT",
  USER = "USER",
}

export interface MessagesDataEvent {
  text: string;
  from: OriginMessage;
  date: string;
}

export interface MessagesData {
  messages: MessagesDataEvent[];
}

export const initialMessages: MessagesData = {
  messages: [],
};

export const messageSlice = createSlice({
  name: "messages",
  initialState: { ...initialMessages },
  reducers: {
    updateMessages: (state, action) => {
      return { ...state, messages: [...state.messages, action.payload] };
    },
    resetMessages: () => {
      return { ...initialMessages };
    },
  },
});


export const {updateMessages, resetMessages} = messageSlice.actions