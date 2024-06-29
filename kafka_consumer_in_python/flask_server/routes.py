from flask import Blueprint, jsonify, request, current_app
from query_handler import QueryHandler

routes_blueprint = Blueprint('routes', __name__)

# Route to send JSON data to Kafka
@routes_blueprint.route('/storeData', methods=['POST'])
def send_to_kafka():
    try:
        handler = QueryHandler()
        handler.insert_data(request.json)
        return jsonify({'success': True, 'message': 'Data sent to Kafka successfully'})
    except Exception as e:
        return jsonify({'success': False, 'error': str(e)})


# Route to start consuming data from Kafka and store in MongoDB
@routes_blueprint.route('/getData', methods=['GET'])
def consume_and_store():
    try:
        handler = QueryHandler(storage=current_app.config["STORAGE_HANDLER"])
        data = handler.query_data(request)
        return jsonify({'success': True, 'message': 'Found the following entries for the search pattern', 'data': data})
    except Exception as e:
        return jsonify({'success': False, 'error': str(e)})

# action.type=crud,actor.id=djkvbwdkjbv