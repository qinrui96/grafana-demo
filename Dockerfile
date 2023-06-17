FROM golang:1.20 as build
WORKDIR /metrics_project
COPY ./src .
RUN go build main.go

FROM alpine as main
COPY --from=build /metrics_project .

CMD ./main