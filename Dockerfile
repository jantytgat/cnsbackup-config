FROM golang:1.17-alpine AS BUILD

WORKDIR $GOPATH/src/github.com/jantytgat/cnsbackup-config
COPY . .

RUN go get -d -v
RUN go install -v


RUN GOOS=linux GOARCH=amd64 go build -o /tmp/cnsbackup-config main.go


FROM gcr.io/distroless/base:latest

COPY --from=BUILD /tmp/cnsbackup-config /usr/local/bin/cnsbackup/cnsbackup-config

CMD ["/usr/local/bin/cnsbackup/cnsbackup-config"]