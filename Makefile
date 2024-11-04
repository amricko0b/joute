.PHONY: clean image

default:
	go mod tidy
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/joute cmd/joute/main.go

image:
	docker build -t amricko0b/joute .

clean:
	rm -rf build