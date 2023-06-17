FROM golang:1.20 as build
WORKDIR /metrics_project
COPY ./src .
RUN ls
RUN go build main.go

FROM ubuntu as main
COPY --from=build /metrics_project/main /bin/
CMD main