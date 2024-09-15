FROM golang:1.22.0 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY main.go ./

RUN go build -o metric-publisher
RUN ls
RUN pwd

FROM alpine:3.20

WORKDIR /app
RUN apk add gcompat
RUN pwd
RUN ls

COPY --from=build /app/metric-publisher metric-publisher 
RUN chmod +x metric-publisher
RUN ls

ENTRYPOINT ["/app/metric-publisher"]