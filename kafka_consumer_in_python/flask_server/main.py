from flask import Flask
from config import STORAGE_STRATEGY, MONGO_DB,MONGO_COLLECTION,MONGO_URI


import sys, os

sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))
from adaptors import mongo_storage

app = Flask(__name__)

import routes
app.register_blueprint(routes.routes_blueprint)

if __name__ == '__main__':
    match STORAGE_STRATEGY:
            case "mongo":
                storage = mongo_storage.MongoStorage(MONGO_URI, MONGO_DB,MONGO_COLLECTION)
                storage.connectToMongo()
                app.config['STORAGE_HANDLER'] = storage
            case _:
                print("UnSupported Storage Strategy")
    app.run(port=8082,debug=True)
