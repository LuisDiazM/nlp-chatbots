import random
import torch
from domain.models.dto.training_model import TrainingModel
from domain.models.dto.models import ModelsDTO
from domain.helpers.constants import COLLECTION_MODELS, COLLECTION_TRAINING, DATABASE_TRAINING
from infraestructure.database.mongoImp import MongoImp
from infraestructure.neural_networks.nn_models.nn_model import NeuralNet
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.storage.s3Imp import S3Imp


class ChatbotResponseUsecase:

    def __init__(self, storage_gateway: S3Imp, preprocessing_nlp: PreprocessingNLP, database_gateway: MongoImp) -> None:
        self.storage_gateway = storage_gateway
        self.preprocessing_nlp = preprocessing_nlp
        self.database_gateway = database_gateway

    def chatbot_response(self, model_id: str, sentence: str):
        model_data = self.__get_model(model_id)
        training_data = self.__get_training_data(model_data.nluIntentId)

        tag = self.__classify_sentence(model_data.modelName, sentence)

        intents = training_data.intents
        for intent in intents:
            if tag == intent.tag:
                max_responses = len(intent.responses)
                random_index = random.randint(0, max_responses-1)
                return intent.responses[random_index]

    def __get_model(self, id: str) -> ModelsDTO:
        model_data = self.database_gateway.find_by_id(
            id, DATABASE_TRAINING, COLLECTION_MODELS)
        return ModelsDTO(**model_data)

    def __get_training_data(self, training_data_id: str):
        intents_data = self.database_gateway.find_by_id(
            training_data_id, DATABASE_TRAINING, COLLECTION_TRAINING)
        return TrainingModel(**intents_data)

    def __classify_sentence(self, model_name: str, sentence: str) -> str:

        model_bin = self.storage_gateway.download(model_name)
        model_data = torch.load(model_bin)

        input_size = model_data["input_size"]
        hidden_size = model_data["hidden_size"]
        output_size = model_data["output_size"]
        all_words = model_data["all_words"]
        tags = model_data["tags"]
        model_state = model_data["model_state"]

        model = NeuralNet(input_size, hidden_size, output_size)
        model.load_state_dict(model_state)
        model.eval()

        tokens = self.preprocessing_nlp.tokenization(sentence)

        X = self.preprocessing_nlp.bag_of_words(tokens, all_words)
        X = X.reshape(1, X.shape[0])
        X = torch.from_numpy(X).to(dtype=torch.float)

        output = model(X)
        _, predicted = torch.max(output, dim=1)
        index = predicted.item()

        return tags[index]
