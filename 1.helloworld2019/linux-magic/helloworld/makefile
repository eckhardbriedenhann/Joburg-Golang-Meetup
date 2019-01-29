.PHONY=run clean


# Builds exe
simplemagic: main.go
	@echo "Building service"
	go build -o $@ .

# Removes exe
clean: 
	@echo "Cleaning service"
	rm simplemagic

# Execute main using runtime
run:
	@echo "Running service"
	go run main.go


