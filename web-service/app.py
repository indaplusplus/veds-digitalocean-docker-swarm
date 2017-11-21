from flask import Flask
import requests

app = Flask(__name__)

@app.route("/")
def index():
    r = requests.get("http://magic-backend-service:7331/")
    return r.text

def main():
    app.run(host="0.0.0.0", port=80)

if __name__ == "__main__":
    main()

