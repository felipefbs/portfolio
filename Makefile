.PHONY: clean build run dev generate templ tailwind install

install:
	go mod download

build:
	go build -o ./tmp/main ./cmd/app/main.go

generate:
	templ generate -path ./templates

tailwind:
	./tailwind -i ./templates/input.css -o ./static/styles/output.css

dev: generate tailwind
	air

run:
	go run ./cmd/app/main.go

clean:
	rm -rf ./tmp
	rm -f ./app
	rm -f tailwind
	rm -f ./static/styles/output.css
	rm -f ./templates/*_templ.go
