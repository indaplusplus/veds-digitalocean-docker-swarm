FROM ubuntu:latest

RUN apt-get update -y
RUN apt-get install -y python-pip python-dev build-essential

ADD . /magic
ENV HOME /magic
WORKDIR /magic

RUN pip install -r requirements.txt

EXPOSE 1337
ENTRYPOINT ["python", "app.py"]
