# Use a multi-stage build
# Build stage
FROM golang:latest AS build


WORKDIR /app


COPY . .


RUN go build -o ./build/app ./cmd/main.go

# Production stage
FROM alpine:latest


WORKDIR /bin

# Copy the binary from the build stage into the production container
COPY --from=build /app .


EXPOSE 3004

RUN chmod +x ./build/app


ENTRYPOINT ["./app"]
