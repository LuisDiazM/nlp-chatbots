import { useFormik } from "formik";
import styles from "./testingChatbot.module.scss";
import { useDispatch, useSelector } from "react-redux";
import {
  MessagesDataEvent,
  OriginMessage,
  updateMessages,
} from "../../../../../redux/states/messages";
import { ACCESS_TOKEN } from "../../../../../shared/utilities/constants";
import { getTestChatbot } from "../../../../utilities/testChatbot";
import SendIcon from '@mui/icons-material/Send';
import { IconButton } from "@mui/material";

interface TestingBotFormProps {
  modelId: string;
}

interface TestingBotFormData {
  modelId: string;
  content: string;
}

const TestingBotForm = ({ modelId }: TestingBotFormProps) => {
  const dispatcher = useDispatch();
  const token = sessionStorage.getItem(ACCESS_TOKEN)??"";

  const userData = useSelector((store: any) => store.user)
  const userId = userData?.email ?? ""
  const initialValuesForm: TestingBotFormData = {
    content: "",
    modelId: modelId,
  };

  const handleSpeakWithBot = async (values: any) => {
    const dateUser = new Date()
    const userMessage: MessagesDataEvent = {
      date: `${dateUser.getHours()}:${dateUser.getSeconds()}`,
      from: OriginMessage.USER,
      text: values.content,
    };
    dispatcher(updateMessages(userMessage));
    const chatResponse = await getTestChatbot(token, values.modelId, values.content, userId)
    const dateBot = new Date()
    const chatbotMessage: MessagesDataEvent = {
      date: `${dateBot.getHours()}:${dateBot.getSeconds()}`,
      from: OriginMessage.BOT,
      text: chatResponse.chat_reponse,
    }
    dispatcher(updateMessages(chatbotMessage));

    formik.setFieldValue("content", "");

  };

  const formik = useFormik({
    initialValues: {
      ...initialValuesForm,
    },
    onSubmit: handleSpeakWithBot,
  });

  return (
    <form className={styles.formContainer} onSubmit={formik.handleSubmit}>
      <input
        id="content"
        type="text"
        placeholder="Escriba aqui para interactuar"
        value={formik.values.content}
        onChange={formik.handleChange}
      />
      <IconButton color="primary" type="submit">
        <SendIcon></SendIcon>
      </IconButton>
    </form>
  );
};

export default TestingBotForm;
