FROM golang:1.15.0-buster

ENV PROTOCFILE protoc-3.13.0-linux-x86_64.zip

ENV PROTOC https://github.com/protocolbuffers/protobuf/releases/download/v3.13.0/${PROTOCFILE}

RUN apt-get update

RUN apt-get install -y unzip

RUN wget ${PROTOC} -O /tmp/${PROTOCFILE}

RUN unzip /tmp/${PROTOCFILE} -d /tmp

RUN cp /tmp/bin/protoc /bin

RUN chmod 0755 /bin/protoc

RUN adduser boxcmd

RUN go get -d -v google.golang.org/protobuf/cmd/protoc-gen-go

RUN go install -v google.golang.org/protobuf/cmd/protoc-gen-go

RUN go get -d -v google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN go install -v google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN go get -d -v google.golang.org/grpc

WORKDIR /go/src/github.com/ronaldoafonso/boxcmd

COPY --chown=boxcmd:boxcmd .ssh/config /home/boxcmd/.ssh/config

COPY --chown=boxcmd:boxcmd ./uci /home/boxcmd/uci

COPY --chown=boxcmd:boxcmd ./gcommand ./gcommand

COPY --chown=boxcmd:boxcmd *.go ./

RUN protoc --go_out=/go/src -I=./gcommand ./gcommand/gcommand.proto

RUN protoc --go-grpc_out=/go/src -I=./gcommand ./gcommand/gcommand.proto

RUN go get -d -v ./... && go install -v ./...

USER boxcmd:boxcmd

CMD ["boxcmd"]
