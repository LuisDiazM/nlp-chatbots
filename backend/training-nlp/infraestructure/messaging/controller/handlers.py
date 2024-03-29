import json
from domain.usecases.notifications_usecase import EventNotificationUsecase

from domain.usecases.storage_models_usecase import StorageModelsUsecase
from cmd.di import Container
from domain.models.train_chatbot_request import DeleteModelsRequest, TrainChatbotRequest
from domain.usecases.training_usecase import TrainingUsecase
from dependency_injector.wiring import Provide, inject
from domain.helpers.loggers import logger
from infraestructure.messaging.controller.constants import FEATURE_TRAININGS, SUBJECT_UPDATE_LICENSE_USAGE 

class ControllerSubscriptions:
    
    @inject
    def __init__(self, training_usecase:TrainingUsecase = Provide[Container.training_usecase],
                 storage_usecase:StorageModelsUsecase=Provide[Container.storage_usecase],
                 notifications_usecase: EventNotificationUsecase=Provide[Container.messaging_client]) -> None:
        self.training_usecase = training_usecase
        self.storage_usecase = storage_usecase
        self.notifications_usecase = notifications_usecase
        super().__init__()

    async def training_model_handler(self, msg):
        subject = msg.subject
        reply = msg.reply
        data = json.loads(msg.data.decode())
        logger.info("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))
        request_data = TrainChatbotRequest(**data)
        model_path = self.training_usecase.generate_model(request_data.id)
        self.storage_usecase.upload_ia_model_to_s3(model_path, request_data.user_id, request_data.id)
        await self.notifications_usecase.publish_notification(SUBJECT_UPDATE_LICENSE_USAGE, {"user_id": request_data.user_id, "feature":FEATURE_TRAININGS})

    
    async def delete_models_by_trainig_id(self, msg):
        subject = msg.subject
        reply = msg.reply
        data = json.loads(msg.data.decode())
        logger.info("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))
        request_data = DeleteModelsRequest(**data)
        training_id = request_data.training_id
        self.storage_usecase.delete_models_by_training_id(training_id)
    
    async def delete_models_by_user_id(self, msg):
        subject = msg.subject
        reply = msg.reply
        data = json.loads(msg.data.decode())
        logger.info("Received a message on '{subject} {reply}': {data}".format(
            subject=subject, reply=reply, data=data))
        user_id = data.get("user_id","")
        if user_id != "":
            self.storage_usecase.delete_models_by_user_id(user_id)