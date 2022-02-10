FROM golang:1.17.6 as builder
WORKDIR /root/gin-mesh
COPY ./ .
RUN CGO_ENABLED=0 GOOS=linux go build -o gin-mesh main.go

FROM alpine:3.15
WORKDIR /bin/
RUN mkdir -p /var/log/gin-mesh
COPY config.yaml .
COPY xds_bootstrap.json .
COPY --from=builder /root/gin-mesh/gin-mesh .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
ENV GRPC_XDS_BOOTSTRAP=./xds_bootstrap.json
ENTRYPOINT [ "/bin/gin-mesh" ]