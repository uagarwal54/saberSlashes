from confluent_kafka import Consumer, KafkaException, KafkaError
import sys
import time

def create_consumer(broker, group, topics):
    conf = {
        'bootstrap.servers': broker,
        'group.id': group,
        'auto.offset.reset': 'earliest',
        'enable.auto.commit': False,
        'fetch.message.max.bytes': 10 * 1024 * 1024,  # 10MB per message
        'fetch.wait.max.ms': 10,  # max wait time for new data
    }
    return Consumer(conf)

def consume_messages(consumer, topics, msg_limit):
    consumer.subscribe(topics)
    msg_count = 0
    start_time = time.time()
    
    try:
        while True:
            msg = consumer.poll(timeout=1.0)
            if msg is None:
                continue
            if msg.error():
                if msg.error().code() == KafkaError._PARTITION_EOF:
                    continue
                elif msg.error():
                    raise KafkaException(msg.error())
            else:
                # Process the message
                print(f'Received message: {msg.value().decode("utf-8")}')
                msg_count += 1
                # Commit offset periodically
                if msg_count % 100 == 0:
                    consumer.commit(asynchronous=True)

                if msg_count >= msg_limit:
                    break

    except KeyboardInterrupt:
        pass
    finally:
        consumer.close()

if __name__ == '__main__':
    broker = 'localhost:9093'
    group = 'high-speed-consumer-group'
    topics = ['high-speed-comm']
    msg_limit = 1000000  # Set to a high value to simulate continuous consumption

    consumer = create_consumer(broker, group, topics)
    consume_messages(consumer, topics, msg_limit)

