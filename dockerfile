FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY conf/conf.yaml /app/conf.yaml
RUN go build -ldflags="-s -w" -o /app/forum main.go

FROM scratch

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/forum /app/forum
COPY --from=builder /app/conf.yaml /app/conf.yaml

CMD ["./forum", "--conf", "conf.yaml"]