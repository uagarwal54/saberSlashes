from pymongo import MongoClient
from common.storage import StorageStrategy

class MongoStorage(StorageStrategy):
    def __init__(self, uri, db_name, collection_name):
        self.client = MongoClient(uri)
        self.db = self.client[db_name]
        self.collection = self.db[collection_name]

    def insert_data(self, data):
        self.collection.insert_one(data)

    def query_data(self, query):
        return list(self.collection.find(query))
