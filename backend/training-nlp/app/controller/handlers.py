import json

from domain.usecases.storage_models_usecase import StorageModelsUsecase
from app.di import Container
from domain.models.train_chatbot_request import TrainChatbotRequest
from domain.usecases.training_usecase import TrainingUsecase
from dependency_injector.wiring import Provide, inject

class ControllerSubscriptions:
    
    @inject
    def __init__(self, training_usecase:TrainingUsecase = Provide[Container.training_usecase], storage_usecase:StorageModelsUsecase=Provide[Container.storage_usecase]) -> None:
        self.training_usecase = training_usecase
        self.storage_usecase = storage_usecase
        super().__init__()

    async def training_model_handler(self, msg):
        subject = msg.subject
        reply = msg.reply
        data = json.loads(msg.data.decode())
        print("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))
        request_data = TrainChatbotRequest(**data)
        model_path = self.training_usecase.generate_model(request_data.id)
        self.storage_usecase.upload_ia_model_to_s3(model_path, request_data.id)