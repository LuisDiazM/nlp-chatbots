import {
  MessagesDataEvent,
  OriginMessage,
} from "../../../../../../redux/states/messages";
import ChatbotMessage from "./chatbotMessages";
import UserMessage from "./userMessages";

export interface MessageEventProps {
  message: MessagesDataEvent;
}

const MessagesFactory = ({ message }: MessageEventProps) => {
  if (message.from === OriginMessage.USER) {
    return <UserMessage message={message}></UserMessage>;
  } else {
    return <ChatbotMessage message={message}></ChatbotMessage>;
  }
};

export default MessagesFactory;
