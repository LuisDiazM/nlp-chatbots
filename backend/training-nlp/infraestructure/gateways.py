from abc import ABC, abstractmethod


class DatabaseGateway(ABC):

    @abstractmethod
    def set_up(self) -> None:
        pass
