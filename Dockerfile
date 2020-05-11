FROM golang:alpine

# Move to working directory app
WORKDIR /app

# Copy code into container
COPY . /app

# Download deps and package required for hit reload
RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

# Build app and run :)
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main -log-prefix=false
