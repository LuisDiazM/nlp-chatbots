import { MessageEventProps } from "./messagesFactory";
import styles from "../testingChatbot.module.scss";

const ChatbotMessage = ({ message }: MessageEventProps) => {
  return (
    <p className={styles.chatbotMessage}>
      {message.text} <span>{message.date}</span>
    </p>
  );
};

export default ChatbotMessage;
