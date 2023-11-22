FROM golang:1.21 AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .
COPY sample-config.yaml .
COPY /internal/db/migration /app/migration

EXPOSE 8080

CMD ["./main", "migrate", "up"]
#CMD echo "Container is running in debug mode" && tail -f /dev/null