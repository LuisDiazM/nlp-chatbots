export interface TrainingChatbotModel {
  userId: string;
  intents: IntentModel[]
  title:string;
}

export interface IntentModel {
  tag: string;
  patterns: string[];
  responses: string[];
}
