import axios from "axios";
import { environment } from "../../environments/environments";

export interface ChatResponse {
  chat_reponse: string;
}

export const getTestChatbot = async (
  token: string,
  modelId: string,
  content: string,
  userId: string
): Promise<ChatResponse> => {
  const URL = environment.BACKEND_URL;
  const response = await axios.get<ChatResponse>(
    `${URL}/nn-response?modelId=${modelId}&content=${content}&userId=${userId}`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
  return response.data;
};
