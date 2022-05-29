
from typing import List
from infraestructure.messaging.nats import NatsImp
from infraestructure.databases.mongo import MongoImp


class HomeUseCase:

    def __init__(self, database: MongoImp, messaging: NatsImp) -> None:
        self.database = database
        self.messaging = messaging
    
    def home(self) -> List:
        return [1, 2, 3, 4]
