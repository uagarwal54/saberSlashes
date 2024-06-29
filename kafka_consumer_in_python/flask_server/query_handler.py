class QueryHandler:
    def __init__(self, storage):
        self.storage = storage

    def query_data(self, request):
        return self.storage.query_data(request)
    
    def insert_data(self, param):
        query = {"field": param}  # Customize your query based on parameters
        return self.storage.insert_data(query)
