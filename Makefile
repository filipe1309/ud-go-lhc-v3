# run tests
test:
	@echo "🟢 Running tests..."

# run node
# example: make run SEC=03
run:
	@echo "🏁 Running code..."
	go run sec-$(SEC)/main.go 

help:
	@echo "📖 Available commands:"
	@echo "  make run SEC=03"
	@echo "  make test"
	@echo "  make help"
