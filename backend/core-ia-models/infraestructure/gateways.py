from abc import ABC, abstractmethod
from typing import Dict


class DatabaseGateway(ABC):

    @abstractmethod
    def set_up(self) -> None:
        pass

    @abstractmethod
    def shutdown(self) -> None:
        pass

    @abstractmethod
    def find_by_id(self, id: str, database: str, collection: str) -> Dict:
        pass


class StorageGateway(ABC):
    

    @abstractmethod
    def download(self, object_name:str) -> None:
        pass

class Messaging(ABC):
    
    @abstractmethod
    async def publish_event(self, subject:str, data:Dict) -> None:
        pass

    @abstractmethod
    async def set_up(self):
        pass

    @abstractmethod
    async def shutdown(self):
        pass