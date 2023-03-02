# Stage 1
FROM golang:alpine as builder

ENV BUILD_DIR /src

WORKDIR ${BUILD_DIR}

# deps
COPY go.mod go.sum ./
RUN  go mod download

# build
COPY . ./
RUN GOOS=linux GOARCH=amd64 GO111MODULE=on CGO_ENABLED=0 \
    go build -o app ./cmd/app/main.go

# Stage 2
FROM alpine:latest

ENV BUILD_DIR /src

COPY --from=builder ${BUILD_DIR} ${BUILD_DIR}

CMD ${BUILD_DIR}/app