from pymongo import MongoClient
from common.storage import StorageStrategy

class MongoStorage(StorageStrategy):
    uri = ""
    db_name = ""
    collection_name = ""
    def __init__(self):
        self.client = MongoClient(self.uri)
        self.db = self.client[self.db_name]
        self.collection = self.db[self.collection_name]

    def insert_data(self, data: dict):
        
        self.collection.insert_one(data)

    def query_data(self, query):
        return list(self.collection.find(query))
