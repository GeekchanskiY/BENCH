FROM golang:1.22
WORKDIR /app

# Mirror
# RUN echo "https://mirror1.alpinelinux.org/alpine/latest-stable/main" > /etc/apk/repositories

RUN apt-get update -y && \
    apt-get install -y --no-cache gcc musl-dev

RUN go install github.com/cosmtrek/air@latest
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

ENV CGO_ENABLED=1
CMD ["air", "-c", ".air.toml"]