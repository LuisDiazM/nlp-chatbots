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