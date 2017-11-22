FROM golang:latest
COPY ./src /go/src/github.com/viktoredstrom/magic-backend-service/backend
WORKDIR /go/src/github.com/viktoredstrom/magic-backend-service/backend
RUN go get ./
RUN go build
EXPOSE 7331
ENTRYPOINT ["backend"]
