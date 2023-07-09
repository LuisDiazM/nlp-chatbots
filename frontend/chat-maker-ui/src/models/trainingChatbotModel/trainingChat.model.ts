export interface TrainingChatbotModel {
  user_id: string;
  intents: IntentModel[]
}

export interface IntentModel {
  tag: string;
  patterns: string[];
  responses: string[];
}
