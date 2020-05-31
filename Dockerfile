# Temel alacağım image golang:alpine.
FROM golang:alpine

# Sonraki dört satırda klasörleri oluşturdum, dosyaları kopyaladım ve çalışma alanını bu klasörü gösterdim.
RUN mkdir /app 
RUN mkdir /app/logs
ADD . /app/
WORKDIR /app 

# Git kurulumunu yaptım ve gerekli kütüphaneleri indirdim. Biraz uzun sürüyor yakında dep ile bunu çözeceğim.
RUN apk update && apk add --no-cache git
RUN go get -v github.com/gin-gonic/gin
RUN go get -v github.com/gorilla/websocket
RUN go get -v go.mongodb.org/mongo-driver/mongo
RUN go get -v github.com/yusufpapurcu/Telemetry

# Tamamını derledim ve Entrypoint ayarladım.
RUN go build -o main .
ENTRYPOINT [ "./main" ]

#docker build -t telempack . -f Dockerfile 
#docker run --name=Telemetry -v /home/yusuftp/log:/app/logs -p 8080:8080 --env-file ./.env telempack -env 