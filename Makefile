.PHONY: all
all: clean compile

APP=message_queue
APP_EXECUTABLE="./out/$(APP)"

# BUILD #############

clean: ##@build remove executable
	go mod tidy -v
	rm -f $(APP_EXECUTABLE)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	go get

test: clean go-get ##@build test the code
	go test ./...

compile: go-get ##@build build the executable
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)