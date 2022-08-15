from infraestructure.database.mongoImp import MongoImp
from helpers.constants import DATABASE_TRAINING, COLLECTION_TRAINING
from domain.models.training_model import TrainingModel


class TrainingUsecase:
    def __init__(self) -> None:
        self.database_gateway = MongoImp()

    def get_training_data(self, id: str) -> TrainingModel:
        data = self.database_gateway.find_by_id(
            id, DATABASE_TRAINING, COLLECTION_TRAINING)
        if data is not None:
            return TrainingModel(**data)
        else:
            return TrainingModel()
