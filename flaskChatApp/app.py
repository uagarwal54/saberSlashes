from flask import Flask, render_template, request, redirect, url_for
from flask_socketio import SocketIO

app = Flask(__name__)
socketio = SocketIO(app)

@app.route('/')
def home():
    return render_template("index.html")


@app.route("/chat")
def chat():
    # This is a dictionary that has all the req params
    username = request.args.get("username")
    room = request.args.get("room")
    if username and room:
        return render_template("chat.html", username=username, room=room)
    else:
        # This will get the  url for the home function
        return redirect(url_for("home"))

# If you open chat.html then there we will find that on connect event which triggers once socketio is able to connect to server,
#  it is emitting a join_room event which the server will recieve
@socketio.on("join_room")
def handle_join_room_event(data):
    app.logger.info("{} someone has joined the room {}".format(data["username"], data["room"]))

if __name__ == "__main__":
    socketio.run(app, debug=True)