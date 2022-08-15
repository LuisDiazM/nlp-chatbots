from infraestructure.database.mongoImp import MongoImp
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.neural_networks.neturalNetImp import NeuralNetImp
from helpers.constants import DATABASE_TRAINING, COLLECTION_TRAINING
from domain.models.training_model import TrainingModel
from domain.models.preprocessing_training_model import PreprocessTrainingModel


class TrainingUsecase:
    def __init__(self) -> None:
        self.database_gateway = MongoImp()
        self.preprocess_nlp = PreprocessingNLP()
        self.neural_net = NeuralNetImp()

    def generate_model(self, id: str) -> str:
        preprocess_data = self.__preprocess_data(id)
        if len(preprocess_data.X_train) > 0 and len(preprocess_data.y_train) > 0:
            return self.neural_net.run_nn(preprocess_data)
        else:
            return ""

    def __preprocess_data(self, id: str) -> PreprocessTrainingModel:
        training_data = self.__get_training_data(id)
        X_train, y_train = self.preprocess_nlp.prepare_data(training_data)
        return PreprocessTrainingModel(
            X_train=X_train, y_train=y_train,
            all_words=self.preprocess_nlp.all_words,
            tags=self.preprocess_nlp.tags, input_size=self.preprocess_nlp.input_len,
            output_size=self.preprocess_nlp.output_len)

    def __get_training_data(self, id: str) -> TrainingModel:
        data = self.database_gateway.find_by_id(
            id, DATABASE_TRAINING, COLLECTION_TRAINING)
        if data is not None:
            return TrainingModel(**data)
        else:
            return TrainingModel()
