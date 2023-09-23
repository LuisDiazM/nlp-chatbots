import { environment } from "../../environments/environments";
import axios from "axios";

export interface TrainingModel {
  user_id: string;
  intents: Intent[];
  title: string;
  description: string;
  id?: string;
}

export interface Intent {
  patterns: string[];
  responses: string[];
  tag: string;
}

export const createTrainingModel = async (
  token: string,
  trainingModel: TrainingModel
) => {
  const URL = environment.BACKEND_URL;
  const response = await axios.post(`${URL}/training-model`, trainingModel, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.data;
};


export const updateTrainingModel = async (
  token: string,
  trainingModel: TrainingModel
) => {
  const URL = environment.BACKEND_URL;
  const data = {...trainingModel}
  delete data.id
  const response = await axios.put(`${URL}/training-model/${trainingModel.id}`, data, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.data;
};

export const deleteTrainingModel = async (
  token: string,
  modelId: string
) => {
  const URL = environment.BACKEND_URL;
  const response = await axios.delete(`${URL}/training-model/${modelId}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.data as TrainingModel[];
};

export const deleteUserAndTrainingModels = async (
  token: string,
  userId: string
) => {
  const URL = environment.BACKEND_URL;
  const response = await axios.delete(`${URL}/training-models/${userId}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.data as TrainingModel[];
};