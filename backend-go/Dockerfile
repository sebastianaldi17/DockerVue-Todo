FROM golang:1.20.1

RUN mkdir -p /go-server
COPY . /go-server

WORKDIR /go-server
RUN go get

RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go mod vendor
EXPOSE 3000


RUN go install -mod=mod github.com/cosmtrek/air@latest
CMD air
# CMD CompileDaemon -log-prefix=false -build="go build" -command="./backend-go"