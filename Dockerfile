FROM golang:alpine AS builder

RUN mkdir /build 
ADD . /build/
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add gcc libc-dev
RUN go get github.com/swaggo/swag/cmd/swag
RUN	go get -u github.com/swaggo/echo-swagger
RUN	swag init -g cmd/main.go
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/main cmd/main.go

FROM scratch
COPY --from=builder /build/bin/main /app/
WORKDIR /app
CMD ["./main"]