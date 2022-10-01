import json

from app.controller.models.train_chatbot_request import TrainChatbotRequest
from domain.usecases.training_usecase import TrainingUsecase
from infraestructure.neural_networks.neturalNetImp import NeuralNetImp
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.database.mongoImp import MongoImp

class ControllerSubscriptions:
    
    def __init__(self, database_gateway:MongoImp, preprocess_nlp:PreprocessingNLP, neural_net:NeuralNetImp) -> None:
        self.database_gateway =database_gateway
        self.preprocess_nlp = preprocess_nlp
        self.neural_net = neural_net

    async def training_model_handler(self, msg):
        subject = msg.subject
        reply = msg.reply
        data = json.loads(msg.data.decode())
        print("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))
        request_data = TrainChatbotRequest(**data)
        usecase = TrainingUsecase(self.database_gateway, self.preprocess_nlp, self.neural_net)
        model_path = usecase.generate_model(request_data.id)
        print(model_path)