FROM golang

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY ./ ./

ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOPROXY=direct
ENV GOPROXY=https://goproxy.io,direct

RUN go mod download

RUN go build -o /server-challenge

EXPOSE 8080

CMD [ "/server-challenge" ]