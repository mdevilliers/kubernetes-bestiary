FROM golang
 
ADD ./src/main.go /go/src/redis-proxy/main.go

RUN cd /go/src/redis-proxy \
		&& go get ./... \
		&& go install redis-proxy
		
ENTRYPOINT /go/bin/redis-proxy

EXPOSE 6379