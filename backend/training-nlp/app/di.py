from domain.usecases.training_usecase import TrainingUsecase
from infraestructure.neural_networks.neturalNetImp import NeuralNetImp
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.database.mongoImp import MongoImp
from dependency_injector import containers, providers


class Container(containers.DeclarativeContainer):

    config = providers.Configuration()

    # infraestructure
    database_client = providers.Singleton(MongoImp)
    preprocess_nlp = providers.Singleton(PreprocessingNLP)
    chatbots_neural_network = providers.Singleton(NeuralNetImp)

    # usecases
    training_usecase = providers.Factory(TrainingUsecase, database_gateway=database_client,
                                         preprocess_nlp=preprocess_nlp, neural_net=chatbots_neural_network)

