import time, json
from datetime import datetime
from confluent_kafka import Producer


jsonTemplate = {
    "time": {},
    "actor":{
        "id": {},
        "group": {}
    },
    "action":{
        "type": {},
        "status": {}
    },
    "resource":{
        "id": {},
        "type": {},
        "status": {} 
    }
}

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
    count = 1
    while time.time() < end_time:
        data = jsonTemplate
        dt_obj = datetime.fromtimestamp(time.time())
        formatted_time = dt_obj.strftime('%Y-%m-%d %H:%M:%S')
        data["time"] = formatted_time
        data["actor"]["id"] = "USR_" + str(count)

        data["actor"]["group"] = "Admin"
        data["action"]["type"] = "CRUD"
        data["action"]["status"] = "success"
        
        data["resource"]["id"] = "RES_" + str(count)
        data["resource"]["type"] = "Lambda"
        data["resource"]["status"] = "Active"
    
        count += 1 

        msg = json.dumps(data)
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
