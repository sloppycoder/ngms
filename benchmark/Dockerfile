FROM golang:1.13 as builder

WORKDIR /build
ADD . /build

RUN go build -o /build/slaver ./slave.go
RUN go build -o /build/client ./client.go

FROM python:3.7-stretch
COPY --from=builder /build/slave /build/client /
RUN pip install locust
ADD ./locustfile.py ./ids.txt ./docker-entrypoint.sh /

ENV RESTSVC_ADDR http://account:3000/accounts
ENV GRPCSVC_ADDR account:3001

EXPOSE 8089

# CMD ["./slave", "-proto", "grpc"]
ENTRYPOINT ["/docker-entrypoint.sh"]