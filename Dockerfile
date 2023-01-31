FROM        golang:alpine
RUN         mkdir -p /app
WORKDIR     /app
COPY        . .
RUN         go mod download && go build -o app
ENTRYPOINT  ["go", "run", "./main.go"]   