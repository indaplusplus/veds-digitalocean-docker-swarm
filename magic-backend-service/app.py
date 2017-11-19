from flask import Flask

app = Flask(__name__)
#
@app.route("/")
def index():
    return "9+10=21"

def main():
    app.run(host="0.0.0.0", port=7331)

if __name__ == "__main__":
    main()

