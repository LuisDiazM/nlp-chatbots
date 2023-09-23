import { useSelector } from "react-redux";
import styles from "../testingChatbot.module.scss";
import { MessagesDataEvent } from "../../../../../../redux/states/messages";
import MessagesFactory from "./messagesFactory";

const MessagesBlock = () => {
  const messages = useSelector((store: any) => store.messages);

  return (
    <div className={styles.messagesBlock}>
      {messages?.messages?.map((message: MessagesDataEvent, index:number) => {
        return <MessagesFactory key={index} message={message}></MessagesFactory>;
      })}
    </div>
  );
};

export default MessagesBlock;
