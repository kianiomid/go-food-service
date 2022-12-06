FROM golang:1.18-alpine as builder

RUN apk add --no-cache git
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /build-food
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o /build-food/cmd/app /build-food/cmd/main.go
RUN ls
FROM alpine:3.14 as production

COPY --from=builder /build-food/cmd/app /app/main
COPY --from=builder /build-food/configurations/App.yaml /app/configurations/App.yaml

WORKDIR /app
RUN chmod +x main
ENTRYPOINT ["./main","app"]
