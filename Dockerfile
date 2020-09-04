FROM golang:1.15.0-buster

RUN adduser boxcmd

WORKDIR /go/src/github.com/ronaldoafonso/boxcmd

COPY --chown=boxcmd:boxcmd .ssh/config /home/boxcmd/.ssh/config

COPY --chown=boxcmd:boxcmd *.go ./

USER boxcmd:boxcmd

RUN go get -d -v ./...

Run go install -v ./...

CMD ["boxcmd"]
