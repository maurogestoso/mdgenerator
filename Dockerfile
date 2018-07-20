FROM golang:alpine
ADD . /go/src/github.com/maurogestoso/mdgenerator
RUN go install -i github.com/maurogestoso/mdgenerator
CMD ["/go/bin/mdgenerator"]
EXPOSE 8080
