FROM golang
 
ADD ./src/main.go /go/src/web-app/main.go

RUN cd /go/src/web-app \
		&& go get ./... \
		&& go install web-app
		
ENTRYPOINT /go/bin/web-app

EXPOSE 8080