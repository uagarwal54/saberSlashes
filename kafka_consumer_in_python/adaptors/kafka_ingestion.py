from kafka import KafkaConsumer
import json
from common.ingestion import IngestionStrategy

class KafkaIngestion(IngestionStrategy):
    grp_id = ""
    topic = ""
    bootstrap_servers = []
    def __init__(self):
        self.consumer = KafkaConsumer(
            self.topic,
            bootstrap_servers=self.bootstrap_servers,
            auto_offset_reset='earliest',
            enable_auto_commit=True,
            group_id=self.grp_id,
            value_deserializer=lambda x: json.loads(x.decode('utf-8'))
        )

    def consume(self):
        for message in self.consumer:
            yield message.value
            # Used yeild as the data will flow in a stream
