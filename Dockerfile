FROM golang:latest
MAINTAINER shanewxy
WORKDIR $GOPATH/src/k8sTraining
ADD . $GOPATH/src/k8sTraining
ENV http_proxy http://172.16.0.92:1087
ENV https_proxy http://172.16.0.92:1087
ENV all_proxy http://172.16.0.92:1087
EXPOSE 9999
RUN go get github.com/mattn/go-sqlite3
RUN go get golang.org/x/crypto/bcrypt
RUN go build -o main
ENTRYPOINT ["./main"]
