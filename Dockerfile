FROM golang:latest
MAINTAINER shanewxy
WORKDIR $GOPATH/src/k8sTraining
ADD . $GOPATH/src/k8sTraining
EXPOSE 9999

RUN go build -o main
ENTRYPOINT ["./main"]
