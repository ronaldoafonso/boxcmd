FROM golang:1.15.0-buster

RUN adduser boxcmd

WORKDIR /go/src/github.com/ronaldoafonso/boxcmd

COPY --chown=boxcmd:boxcmd .ssh /home/boxcmd/.ssh

COPY --chown=boxcmd:boxcmd main.go .

USER boxcmd:boxcmd

RUN go get -d -v ./...

Run go install -v ./...

CMD ["boxcmd"]
