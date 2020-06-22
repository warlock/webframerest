hello: echo "Hello webframerest"

build: go build -ldflags "-s -w"

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

doc: 
	swag init

all: hello doc build