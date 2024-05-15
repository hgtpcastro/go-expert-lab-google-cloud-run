# Build Stage
FROM golang:1.21 as BuildStage
WORKDIR /app
COPY . .
# RUN go mod download
EXPOSE 8080
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o weather-zip-code-api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./deployments/bin/weather-zip-code-api ./internal/services/zip_code_service/cmd/app/main.go

# Deploy Stage
FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage ./deployments/bin/weather-zip-code-api .
EXPOSE 8080
ENTRYPOINT [ "/app/weather-zip-code-api"]