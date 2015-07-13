FROM golang
 
ADD ./src/main.go /go/src/static_web_server/main.go
RUN go install static_web_server
ENTRYPOINT /go/bin/static_web_server

EXPOSE 8080