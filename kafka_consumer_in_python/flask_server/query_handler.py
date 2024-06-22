from common.storage import StorageStrategy

class QueryHandler:
    def __init__(self, storage: StorageStrategy):
        self.storage = storage

    def query_data(self, param):
        query = {"field": param}  # Customize your query based on parameters
        return self.storage.query_data(query)
