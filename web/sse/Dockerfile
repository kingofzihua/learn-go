FROM golang:1.20 as builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn 

RUN mkdir -p bin/ && go build -o ./bin/ ./...

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 80

VOLUME /data/conf

CMD ["./sse"]