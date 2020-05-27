FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN apk update && apk add --no-cache git
RUN go get -v github.com/gin-gonic/gin
RUN go get -v github.com/gorilla/websocket
RUN go get -v go.mongodb.org/mongo-driver/mongo
RUN go get -v github.com/yusufpapurcu/Telemetry
RUN go build -o main .
ENTRYPOINT [ "./main" ]