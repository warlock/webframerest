FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app
RUN apk add gcc libc-dev
RUN go get github.com/swaggo/swag/cmd/swag
RUN	go get -u github.com/swaggo/echo-swagger
RUN	swag init
RUN go build -ldflags "-s -w"
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./webframerest"]