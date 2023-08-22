FROM golang:1.19

ARG envType=test

WORKDIR /usr/src/app

COPY . .

COPY configs/config_${envType}.toml ./configs/config.toml

RUN go env && go env -w GOPROXY=https://goproxy.cn,direct && go build -o ./bin/main && chmod +x ./bin/main

CMD ["./bin/main", "-c", "./configs/config.toml"]

EXPOSE 8080