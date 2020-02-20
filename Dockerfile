FROM golang:1.13-stretch

WORKDIR /go/src/github.com/lucdecas/dockerdevhelper

COPY . .

RUN go build -o /usr/bin/dockerdevhelper dockerdevhelper.go

EXPOSE "8081"

CMD ["dockerdevhelper"]