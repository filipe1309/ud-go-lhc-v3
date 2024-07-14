# run section test
test:
	@echo "🟢 Running tests..."
	go test -v ./sec-$(SEC)/...

# run all tests
test-all:
	@echo "🟢 Running all tests..."
	go test -v ./...

build:
	@echo "🔨 Building code..."
	go build sec-$(SEC)/main.go

# run node
# example: make run SEC=03
run:
	@echo "🏁 Running code..."
	go run sec-$(SEC)/main.go 

# create new section using scripts/new-section.sh
# example: make new-section SEC=03
new-section:
	@echo "📝 Creating new section $(SEC)..."
	./scripts/new-section.sh $(SEC)
	

help:
	@echo "📖 Available commands:"
	@echo "  make run SEC=03"
	@echo "  make test"
	@echo "  make help"
