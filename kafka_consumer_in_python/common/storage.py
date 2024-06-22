from abc import ABC, abstractmethod

class StorageStrategy(ABC):
    @abstractmethod
    def insert_data(self, data):
        pass

    @abstractmethod
    def query_data(self, query):
        pass
