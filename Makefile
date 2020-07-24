build:
	go build -ldflags "-s -w" -o bin/main main.go

freebsd:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

linux:
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go

windows:
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

doc: 
	go get github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/echo-swagger
	swag init

all: doc build