FROM golang:1.4
MAINTAINER tobe tobeg3oogle@gmail.com

ADD . /go/src/

WORKDIR /go/src/

CMD /bin/bash