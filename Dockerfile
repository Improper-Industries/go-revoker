FROM golang:alpine3.16

RUN mkdir /revoker
WORKDIR /revoker
COPY . .

RUN GOPROXY=https://goproxy.io go get -d -v ./...
RUN go build .

EXPOSE 3005

CMD "./revoker"
