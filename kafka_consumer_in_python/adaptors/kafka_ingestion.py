from confluent_kafka import Consumer, KafkaException, KafkaError
from common.ingestion import IngestionStrategy
import json

class KafkaIngestion(IngestionStrategy):
    grp_id = ""
    topics = []
    bootstrap_servers = ""
    def __init__(self):
        conf = {
            'bootstrap.servers': self.bootstrap_servers,
            'group.id': self.grp_id,
            'auto.offset.reset': 'earliest',
            'enable.auto.commit': False,
            'fetch.message.max.bytes': 10 * 1024 * 1024,  # 10MB per message
            'fetch.wait.max.ms': 10,  # max wait time for new data
        }
        self.consumer = Consumer(conf)
        self.consumer.subscribe(topics=self.topics)

    def consume(self):
        # for message in self.consumer:
        #     yield message.value
        #     # Used yeild as the data will flow in a stream
        try:
            while True:
                msg = self.consumer.poll(timeout=1.0)
                if msg is None:
                    continue
                if msg.error():
                    if msg.error().code() == KafkaError._PARTITION_EOF:
                        continue
                    elif msg.error():
                        raise KafkaException(msg.error())
                else:
                    # Process the msg
                    m = json.loads(msg.value().decode("utf-8"))
                    yield m

        except KeyboardInterrupt:
            pass
        finally:
            self.consumer.close()