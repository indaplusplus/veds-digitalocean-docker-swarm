from flask import Flask, render_template
import requests
import json

app = Flask(__name__)

def api_request(site, request_type):
    r = requests.get("http://magic-backend-service:7331/%s/%s" % (site, request_type))
    return r.text

@app.route("/")
def index():
    return "try out our amazing services @ /dn/ and /svd/"

@app.route("/dn/")
def dn():
    return news("dn")

@app.route("/svd/")
def svd():
    return news("svd")

def news(site):
    newest_article = json.loads(api_request(site, "recent"))
    articles = json.loads(api_request(site, "all"))
    return render_template("news.html", recent = newest_article, all_articles = articles)

def main():
    app.run(host="0.0.0.0", port=80)

if __name__ == "__main__":
    main()

