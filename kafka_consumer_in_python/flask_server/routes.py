from flask import Blueprint, jsonify, request
from query_handler import QueryHandler

routes_blueprint = Blueprint('routes', __name__)

# Route to send JSON data to Kafka
@routes_blueprint.route('/send_to_kafka', methods=['POST'])
def send_to_kafka():
    try:
        handler = QueryHandler()
        handler.insert_data(request.json)
        return jsonify({'success': True, 'message': 'Data sent to Kafka successfully'})
    except Exception as e:
        return jsonify({'success': False, 'error': str(e)})


# Route to start consuming data from Kafka and store in MongoDB
@routes_blueprint.route('/consume_and_store', methods=['GET'])
def consume_and_store():
    try:
        handler = QueryHandler()
        handler.query_data(request.json)
        return jsonify({'success': True, 'message': 'Started consuming data from Kafka and storing in MongoDB'})
    except Exception as e:
        return jsonify({'success': False, 'error': str(e)})
