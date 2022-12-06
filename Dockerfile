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
# Add current directory
#WORKDIR /app/food/cmd
RUN ls
# Build app
#RUN go build -o /build/cmd/http/app /build/cmd/http/main.go
RUN go build -o /build-food/cmd/app /build-food/cmd/main.go
#RUN go build -o . app/food
#RUN go build -a -o ./app/food
RUN ls
FROM alpine:3.14 as production
# Copy built binary from builder
#COPY --from=builder app/food .

#COPY --from=build-env /build/cmd/http/app /app/main
COPY --from=builder /build-food/cmd/app /app/main
#COPY --from=build-env /build/cmd/http/config.json /app/config.json
COPY --from=builder /build-food/configurations/App.yaml /app/configurations/App.yaml

#COPY --from=builder /out /food
#COPY --from=build /app/ht/src/out/ht /ht
RUN ls

WORKDIR /app
RUN chmod +x main
ENTRYPOINT ["./main","app"]

#RUN chmod +x app/food
# Expose port
#EXPOSE 8090
# Exec built binary
#CMD ["./app/food"]