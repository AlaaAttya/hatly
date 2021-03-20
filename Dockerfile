FROM golang

WORKDIR /src

RUN go get -u github.com/yyh-gl/realize
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

COPY ./src /app
WORKDIR /app

RUN go mod download

EXPOSE 80

CMD ["/wait", "go run main.go"]