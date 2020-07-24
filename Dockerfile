FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN make all
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./bin/main"]