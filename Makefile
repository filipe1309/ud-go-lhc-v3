# run tests
test:
	@echo "🟢 Running tests..."

# run node
run:
	@echo "🏁 Running code..."
	go run sec-03/exercise_1.go 

help:
	@echo "📖 Available commands:"
	@echo "  make run"
	@echo "  make test"
	@echo "  make help"
