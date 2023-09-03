from domain.usecases.gateways import DatabaseGateway
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.neural_networks.neturalNetImp import NeuralNetImp
from domain.helpers.constants import DATABASE_TRAINING, COLLECTION_TRAINING
from domain.models.dto.training_model import TrainingModel
from domain.models.preprocessing_training_model import PreprocessTrainingModel


class TrainingUsecase:
    def __init__(self, database_gateway: DatabaseGateway, preprocess_nlp: PreprocessingNLP, neural_net: NeuralNetImp) -> None:
        self.database_gateway = database_gateway
        self.preprocess_nlp = preprocess_nlp
        self.neural_net = neural_net

    def generate_model(self, id: str) -> str:
        preprocess_data = self.__preprocess_data(id)
        if len(preprocess_data.x_train) > 0 and len(preprocess_data.y_train) > 0:
            return self.neural_net.run_nn(preprocess_data)
        else:
            return ""

    def __preprocess_data(self, id: str) -> PreprocessTrainingModel:
        training_data = self.__get_training_data(id)
        x_train, y_train = self.preprocess_nlp.prepare_data(training_data)
        return PreprocessTrainingModel(
            x_train=x_train, y_train=y_train,
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
