include .env

# targets that do not relate to files
.PHONY=run docker-build docker-run

serviceDetails="$(name) version:$(version)"
dockertag=$(name):$(version)

run: $(deps)
	@echo "Running service: $(serviceDetails)"
	go run main.go

$(name): $(deps)
	@echo "Building service: $(serviceDetails) "
	go build -o $@ .

clean: 
	@echo "Cleaning service: $(serviceDetails) "
	rm $(name)

docker-build:
	@echo "Building Docker image service: $(serviceDetails) "
	docker build -t $(dockertag) .

docker-run: docker-build
	docker run -it --rm -p 8081:8081 $(dockertag) magic
