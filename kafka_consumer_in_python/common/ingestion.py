from abc import ABC, abstractmethod

class IngestionStrategy(ABC):
    @abstractmethod
    def consume(self):
        pass
