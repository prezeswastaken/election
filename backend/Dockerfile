FROM golang:1.21.3-bookworm
WORKDIR /app
COPY . .

# Install go dependencies
RUN go mod tidy
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
