
BINARY_NAME=food-delivery

## build: Build binary
build:
	@echo "Building back end..."
	go build -o ${BINARY_NAME} 
	@echo "Binary built!"

## run: builds and runs the application
run: build
	@echo "Starting back end..."
	@env DSN=${DSN} ./${BINARY_NAME} &
	@echo "Back end started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping back end..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped back end!"

## restart: stops and starts the running application
restart: stop start

protoc:
	protoc -I ./proto \
	--go_out ./pb --go_opt paths=source_relative \
	--go-grpc_out ./pb --go-grpc_opt paths=source_relative \
	./proto/demo.proto
	