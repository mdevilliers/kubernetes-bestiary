FROM golang
 
ADD ./src/main.go /go/src/basic_web_server/main.go
RUN go install basic_web_server
ENTRYPOINT /go/bin/basic_web_server

EXPOSE 8080