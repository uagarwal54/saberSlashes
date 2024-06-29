from pymongo import MongoClient
from common.storage import StorageStrategy

class MongoStorage(StorageStrategy):
    uri = ""
    db_name = ""
    collection_name = ""
    def __init__(self, uri, dbName, collName):
        self.uri = uri
        self.db_name = dbName
        self.collection_name = collName

    def connectToMongo(self):
        print("Trying to connect to mongo")
        self.client = MongoClient(self.uri)
        self.db = self.client[self.db_name]
        self.collection = self.db[self.collection_name]
        print("Connected to mongo")
        
    def insert_data(self, data: dict):
        self.collection.insert_one(data)

    def query_data(self, request):
        searchParams = list(request.args.keys())
        query = {}
        for key in searchParams:
            query[key] = request.args.get(key)
        return list(self.collection.find(query, {'_id': 0}))
      # We have used {'_id': 0} so that the resulting data will NOT contain the _id field
