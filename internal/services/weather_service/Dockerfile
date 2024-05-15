# Build Stage
FROM golang:1.21-alpine as builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /services/weather_service

# Install dependencies
RUN apk --update --no-cache add ca-certificates

# Build Go binary
COPY go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.io,direct/
RUN go mod download
COPY . .
RUN go mod tidy

# Skaffold passes in debug-oriented compiler flags
ARG SKAFFOLD_GO_GCFLAGS
RUN go build.sh -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /services/weather_service .

# Deployment container
FROM scratch

WORKDIR /weather_service

# Definition of this variable is used by 'skaffold debug' to identify a golang binary.
# Default behavior - a failure prints a stack trace for the current goroutine.
# See https://golang.org/pkg/runtime/
ENV GOTRACEBACK=single

COPY --from=builder /services/weather_service /app/weather-api

ENTRYPOINT ["/app/weather-api"]