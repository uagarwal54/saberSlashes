INGESTOR_STRATEGY = "kafka"
STORAGE_STRATEGY = "mongo"

KAFKA_TOPIC = ["high-speed-comm"]
KAFKA_BOOTSTRAP_SERVERS = 'localhost:9093'
KAFKA_GROUP_ID = "demo"


MONGO_URI = "mongodb://root:example@mongo:27017/"
MONGO_DB = "consumer_logs"
MONGO_COLLECTION = "logs"