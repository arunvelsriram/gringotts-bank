SERVICES = frontend recommendation customer
BUILD_SERVICE_TARGETS=$(foreach service,$(SERVICES),build/$(service))
RUN_SERVICE_TARGETS=$(foreach service,$(SERVICES),run/$(service))

build: clean $(BUILD_SERVICE_TARGETS)

deps:
	@go mod tidy -v
	@go mod vendor

clean:
	@rm -rf bin

build/%:
	@echo "Building $(@F)..."
	@go build -o bin/$(@F) ./cmd/$(@F)

run: build
	@echo "Running all services..."
	@parallel --line-buffer ::: $(foreach service,$(SERVICES),./bin/$(service))

run/%:
	@echo "Running $(@F)..."
	@go run ./cmd/$(@F)/main.go
