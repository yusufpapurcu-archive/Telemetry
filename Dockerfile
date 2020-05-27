FROM golang:alpine
RUN mkdir /app 
RUN mkdir /app/logs
ADD . /app/
WORKDIR /app 
RUN apk update && apk add --no-cache git
RUN go get -v github.com/gin-gonic/gin
RUN go get -v github.com/gorilla/websocket
RUN go get -v go.mongodb.org/mongo-driver/mongo
RUN go get -v github.com/yusufpapurcu/Telemetry
RUN go build -o main .
ENTRYPOINT [ "./main" ]
#sudo docker build -t telempack . -f Dockerfile 
#sudo docker run --name=Telemetry -v /home/yusuftp/log:/app/logs -p 8080:8080 --env-file ./.env telempack -env