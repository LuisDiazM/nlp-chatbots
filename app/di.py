from dependency_injector import containers, providers
from domain.usecases.chatbot_response_usecase import ChatbotUsecase
from domain.usecases.home_usecase import HomeUseCase
from infraestructure.messaging.nats import NatsImp
from infraestructure.databases.mongo import MongoImp


class Container(containers.DeclarativeContainer):
    # infraestructure
    config = providers.Configuration()
    database = providers.Singleton(MongoImp)
    messaging = providers.Singleton(NatsImp)

    # usecases
    home_usecase = providers.Factory(
        HomeUseCase, database=database, messaging=messaging)
