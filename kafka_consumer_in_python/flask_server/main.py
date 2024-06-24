from flask import Flask
from routes import routes_blueprint

app = Flask(__name__)
app.register_blueprint(routes_blueprint)

if __name__ == '__main__':
    app.run(port=8080,debug=True)

# https://chatgpt.com/share/e53f3144-e4c6-4c3e-956d-0a513a974d95