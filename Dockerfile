FROM golang

WORKDIR /src

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

COPY ./ /src

RUN go mod download

EXPOSE 80

CMD ["/wait", "go run main.go"]