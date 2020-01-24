APP=newrelic
PORT=4000
CLIENTS=5
CLIENT_PATH=./cmd/client/main.go
SERVER_PATH=./cmd/server/main.go

#Server
.PHONY: build-server
build-server:
				go build -o ${APP} ${SERVER_PATH}

.PHONY: run-server
run-server:
				go run -race ${SERVER_PATH} --port ${PORT} --clients ${CLIENTS}

.PHONY: test
test:
				go test -v ./...

#Client
.PHONY: build-client
build-client:
				go build -o ${APP} ${CLIENT_PATH}

.PHONY: run-client
run-client:
				go run -race ${CLIENT_PATH} --port ${PORT}

.PHONY: clean
clean:
				go clean

.PHONY: docker-server-run
docker-server-run:
			docker build --build-arg path=${SERVER_PATH} -t ${APP} .
			docker run -e path=:${SERVER_PATH} ${APP}

.PHONY: help
## help: Print this help message
## build-server: Build the server
## build-client: Build sthe client
## run-server: Run the server with arguments --port ${PORT} --clients ${CLIENTS}
## run-client: Run the client with arguments --port ${PORT}
## docker-server-run: Build the image and run the container for the server
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

