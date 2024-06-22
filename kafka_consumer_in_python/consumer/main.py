from adaptors.kafka_ingestion import KafkaIngestion
from adaptors.mongo_storage import MongoStorage
from common.ingestion import IngestionStrategy
from common.storage import StorageStrategy
from config import KAFKA_TOPIC, KAFKA_BOOTSTRAP_SERVERS, MONGO_URI, MONGO_DB, MONGO_COLLECTION

class ConsumerService:
    def __init__(self, ingester: IngestionStrategy, storage: StorageStrategy):
        self.ingester = ingester
        self.storage = storage

    def run(self):
        for data in self.ingester.consume():
            self.storage.insert_data(data)

if __name__ == '__main__':
    ingester = KafkaIngestion(KAFKA_TOPIC, KAFKA_BOOTSTRAP_SERVERS)
    storage = MongoStorage(MONGO_URI, MONGO_DB, MONGO_COLLECTION)
    service = ConsumerService(ingester, storage)
    service.run()
