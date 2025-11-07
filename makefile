BINARY = modbus-server

MAIN = ./cmd/modbus-server

CONFIG = config.yaml

MODBUS_PORT = 5020
API_PORT = 8080

# --------------------------
# GO Commands
# --------------------------
.PHONY: all build run clean docker-build docker-run

all: build

build:
	@echo "Build the binary..."
	go mod tidy
	go build -o $(BINARY) $(MAIN)

run: build
	@echo "Server launching..."
	./$(BINARY) -config $(CONFIG)

clean:
	@echo "Cleaning..."
	rm -f $(BINARY)

# --------------------------
# Docker
# --------------------------
docker-build:
	@echo "Docker image build..."
	docker build -t $(BINARY) .

docker-run:
	@echo "Launching the Docker container..."
	docker run -p $(MODBUS_PORT):$(MODBUS_PORT) -p $(API_PORT):$(API_PORT) \
		-v $(PWD)/$(CONFIG):/app/$(CONFIG) \
		$(BINARY)
