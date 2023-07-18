# arch-sample-service dockerfile

# stage 1 - build
FROM golang:1.20-alpine AS build

COPY ./cmd /arch-sample-cli/cmd

WORKDIR /arch-sample-cli

COPY go.mod ./
COPY go.sum ./

RUN go get -d -v ./...
RUN go install -v ./...

COPY *.go ./

RUN go build -o bin/arch-sample-cli

EXPOSE 8080

# runs forever, so terminal into the host 
ENTRYPOINT exec top -b

# executes the CLI with arguments and exits the container after done (cannot shell in)
# ENTRYPOINT ["./bin/arch-sample-cli", "--version"]