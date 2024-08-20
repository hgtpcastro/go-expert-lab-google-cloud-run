# Start from golang base image
FROM golang:1.21-alpine as builder

# Install Bash inside the container
RUN apk update && apk add bash

# Set the current working directory inside the container
WORKDIR /go-expert-lab-google-cloud-run/weather_service

# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy sources to the working directory
COPY . .

# Set the Go environment
ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOARCH amd64

# Set the App environment
ARG app_env
ENV APP_ENV $app_env

# Run the app
ARG project
ENV PROJECT $project

CMD go run ${PROJECT}
# ENTRYPOINT ["tail", "-f", "/dev/null"]