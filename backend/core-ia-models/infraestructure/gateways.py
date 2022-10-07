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

    @abstractmethod
    def insert_one(self, database:str, collection:str, data:Dict) -> None:
        pass


class StorageGateway(ABC):
    
    @abstractmethod
    def upload(self, path_file: str) -> str:
        pass

    @abstractmethod
    def download(self, object_name:str) -> None:
        pass
