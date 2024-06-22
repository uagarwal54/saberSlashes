from adaptors.kafka_ingestion import KafkaIngestion
from adaptors.mongo_storage import MongoStorage
from common.ingestion import IngestionStrategy
from common.storage import StorageStrategy
from config import INGESTOR_STRATEGY, STORAGE_STRATEGY,KAFKA_TOPIC, KAFKA_BOOTSTRAP_SERVERS, KAFKA_GROUP_ID, MONGO_URI, MONGO_DB, MONGO_COLLECTION

class ConsumerService:
    def __init__(self, ingester: IngestionStrategy, storage: StorageStrategy):
        self.ingester = ingester
        self.storage = storage

    def run(self):
        for data in self.ingester.consume():
            self.storage.insert_data(data)

if __name__ == '__main__':
    ingester = ""
    storage = ""
    match INGESTOR_STRATEGY:
        case "kafka":
            KafkaIngestion.bootstrap_servers = KAFKA_BOOTSTRAP_SERVERS
            KafkaIngestion.topics = KAFKA_TOPIC
            KafkaIngestion.grp_id = KAFKA_GROUP_ID
            ingester = KafkaIngestion()
        case _:
            print("The defined value for INGESTOR_STRATEGY is not supported")
    
    match STORAGE_STRATEGY:
        case "mongo":
            MongoStorage.db_name = MONGO_DB
            MongoStorage.collection_name = MONGO_COLLECTION
            MongoStorage.uri = MONGO_URI
            storage = MongoStorage()
        case _:
            print("The defined value for STORAGE_STRATEGY is not supported")

    
    service = ConsumerService(ingester, storage)
    service.run()
