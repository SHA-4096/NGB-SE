FROM golang:1.18-alpine3.15 as builder
COPY ./ /src
WORKDIR /src
ENV GOPROXY "https://goproxy.cn"
RUN go build -o /build/app cmd/server.go

FROM alpine:3.15
COPY --from=builder /build/app /usr/bin/app
COPY --from=builder /src/config/config.yaml /config/config.yaml
WORKDIR /
ENTRYPOINT [ "app" ]