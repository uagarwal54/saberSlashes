from flask import jsonify, request
from app import app, producer, start_consuming

# Route to send JSON data to Kafka
@app.route('/send_to_kafka', methods=['POST'])
def send_to_kafka():
    try:
        data = request.json
        producer.send(KAFKA_TOPIC, data)
        return jsonify({'success': True, 'message': 'Data sent to Kafka successfully'})
    except Exception as e:
        return jsonify({'success': False, 'error': str(e)})


# Route to start consuming data from Kafka and store in MongoDB
@app.route('/consume_and_store', methods=['GET'])
def consume_and_store():
    try:
        start_consuming()
        return jsonify({'success': True, 'message': 'Started consuming data from Kafka and storing in MongoDB'})
    except Exception as e:
        return jsonify({'success': False, 'error': str(e)})
