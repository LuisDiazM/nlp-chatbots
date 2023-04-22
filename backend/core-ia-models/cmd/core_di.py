from domain.usecases.chatbot_responses_usecase import ChatbotResponseUsecase
from infraestructure.storage.s3Imp import S3Imp
from infraestructure.preprocessing.nlp_preprocessing import PreprocessingNLP
from infraestructure.database.mongoImp import MongoImp
from dependency_injector import containers, providers


class Container(containers.DeclarativeContainer):

    config = providers.Configuration()

    # infraestructure
    database_client = providers.Singleton(MongoImp)
    preprocess_nlp = providers.Singleton(PreprocessingNLP)
    storage_client = providers.Singleton(S3Imp)

    # usecases
    chatbot_usecase = providers.Factory(ChatbotResponseUsecase, storage_gateway=storage_client, preprocessing_nlp=preprocess_nlp, database_gateway=database_client)

