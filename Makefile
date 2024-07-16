# run section test
test:
	@echo "🟢 Running tests..."
	@if [ -z "$(SEC)" ]; then \
		SEC=$$(ls -d sec-* | tail -n 1 | cut -d '-' -f 2); \
		echo "🔍 No section provided, running last section: $$SEC"; \
		go test -v sec-$$SEC/...; \
	else \
		go test -v sec-$(SEC)/...; \
	fi

# run all tests
test-all:
	@echo "🟢 Running all tests..."
	go test -v ./...

build:
	@echo "🔨 Building code..."
	@if [ -z "$(SEC)" ]; then \
		SEC=$$(ls -d sec-* | tail -n 1 | cut -d '-' -f 2); \
		echo "🔍 No section provided, running last section: $$SEC"; \
		go build sec-$$SEC/main.go; \
	else \
		go build sec-$(SEC)/main.go \
	fi
	

# run node
# example: make run SEC=03
run:
	@echo "🏁 Running code..."
	@if [ -z "$(SEC)" ]; then \
		SEC=$$(ls -d sec-* | tail -n 1 | cut -d '-' -f 2); \
		echo "🔍 No section provided, running last section: $$SEC"; \
		go run sec-$$SEC/main.go; \
	else \
		go run sec-$(SEC)/main.go; \
	fi

# create new section using scripts/new-section.sh
# example: make new-section [SEC=03] # default last section + 1
new-section:
	@echo "📝 Creating new section $(SEC)..."
	./scripts/new-section.sh $(SEC)
	

help:
	@echo "📖 Available commands:"
	@echo "  make run SEC=03"
	@echo "  make test SEC=03"
	@echo "  make test-all"
	@echo "  make build"
	@echo "  make new-section [SEC=03] # default last section + 1"
	@echo "  make help"
