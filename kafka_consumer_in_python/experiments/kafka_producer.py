import time
from confluent_kafka import Producer

def acked(err, msg):
    if err is not None:
        print(f"Failed to deliver message: {err}")
    else:
        print(f"Message produced: {msg.value().decode('utf-8')}")

def create_producer(broker):
    conf = {
        'bootstrap.servers': broker,
        'linger.ms': 10,  # Small delay to batch messages
        'batch.num.messages': 1000,  # Number of messages to batch
    }
    return Producer(conf)

def produce_messages(producer, topic, msg_rate, duration):
    msg_count = 0
    start_time = time.time()
    end_time = start_time + duration
    interval = 1.0 / msg_rate

    while time.time() < end_time:
        msg = f"Message {msg_count}"
        producer.produce(topic, msg.encode('utf-8'), callback=acked)
        msg_count += 1
        producer.poll(0)  # Serve delivery reports

        # Wait to maintain the message rate
        time.sleep(interval)

    # Wait for all messages to be delivered
    producer.flush()

if __name__ == '__main__':
    broker = 'localhost:9093'
    topic = 'high-speed-comm'
    msg_rate = 1000  # Messages per second
    duration = 15  # Duration in seconds
    producer = create_producer(broker)
    produce_messages(producer, topic, msg_rate, duration)
