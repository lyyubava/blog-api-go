FROM golang:1.20 AS build

RUN useradd -u 1001 -m gouser
WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

FROM alpine:latest
RUN apk add --no-cache bash
WORKDIR /

COPY --from=build /main /main

USER 1001
EXPOSE 8080

ENTRYPOINT ["/main"]