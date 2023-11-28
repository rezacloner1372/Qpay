# Stage 1: Build the application
FROM golang:1.21 AS build

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /Qpay .

# Stage 2: Create the final lightweight image
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash

WORKDIR /app

COPY --from=build /Qpay .
COPY sample-config.yaml .
COPY internal/db/migration /app/internal/db/migration
COPY wait-for-it.sh /app/wait-for-it.sh 
RUN chmod +x /app/wait-for-it.sh

EXPOSE 8080

# ENTRYPOINT [ "/bin/bash", "-c" ]
# CMD ["./Qpay"]