FROM golang:1.22-alpine
 
WORKDIR /app

RUN apk update && \
    apk add gcc musl-dev

RUN go install github.com/cosmtrek/air@latest
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

ENV CGO_ENABLED=1
CMD ["air", "-c", ".air.toml"]