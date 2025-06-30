FROM golang:1.22.0 AS buildbase

WORKDIR /usr/src/app

COPY go.mod ./

#RUN go mod download

COPY . .

RUN go build -o image-api

EXPOSE 7999

CMD ["./image-api"]
#RUN CGO_ENABLED=0 GOOS=linux go build -o image-api ./image-api
