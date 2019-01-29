include .env

# targets that do not relate to files
.PHONY=build run

serviceDetails="$(name) version:$(version)"

run: $(deps)
	@echo "Running service: $(serviceDetails)"
	go run main.go

$(name): $(deps)
	@echo "Building service: $(serviceDetails) "
	go build -o $@ .

clean: 
	@echo "Cleaning service: $(serviceDetails) "
	rm $(name)