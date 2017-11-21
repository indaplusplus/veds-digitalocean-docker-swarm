docker login --username veds

docker build -t magic-backend-service ./magic-backend-service
docker tag magic-backend-service veds/homework-kth-backend
docker push veds/homework-kth-backend

docker build -t web-service ./web-service
docker tag web-service veds/homework-kth-webservice
docker push veds/homework-kth-webservice
