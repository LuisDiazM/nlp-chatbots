import { MessageEventProps } from "./messagesFactory";
import styles from "../testingChatbot.module.scss";

const UserMessage = ({ message }: MessageEventProps) => {
  return (
    <p className={styles.userMessage}>
      {message.text} <span>{message.date}</span>
    </p>
  );
};

export default UserMessage;
