run:
	go run cmd/main.go

build:
	go build -ldflags "-s -w" -o bin/main cmd/main.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/main cmd/main.go

linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/main cmd/main.go

windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/main cmd/main.go

doc: 
	go get github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/echo-swagger
	swag init

all: doc build