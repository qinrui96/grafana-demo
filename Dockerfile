FROM golang:1.20
WORKDIR /metrics_project
COPY . .
RUN go build src/main.go

CMD ./main