# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------
FROM golang:1.16.8-alpine3.14 AS builder
ENV GO111MODULE=on

# Set working directory
WORKDIR /go/src/github.com/ntsd/alpha-interface-backend

# install packr
RUN go get -u github.com/gobuffalo/packr/packr

# Download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy stuff from development stage
COPY . .

# Build the binary
RUN packr build -o /go/bin/server.bin ./cmd/main.go

# ------------------------------------------------------------------------------
# Application image
# ------------------------------------------------------------------------------
FROM golang:1.15.4-alpine3.12 as release

# Set working directory
WORKDIR /app

RUN apk add --update --no-cache ca-certificates bash

RUN addgroup -g 1000 appuser && \
  adduser -u 1000 -G appuser -g appuser -s /bin/sh -D appuser && \
  chown 1000:1000 /app

#Get artifact from builder stage
COPY --chown=1000:1000 --from=builder /go/bin/server.bin /app/server.bin

USER appuser

EXPOSE 8001

CMD ["sh", "-c", "/app/server.bin"]
