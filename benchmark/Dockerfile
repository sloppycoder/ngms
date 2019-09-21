FROM golang:1.13 as builder

WORKDIR /build
ADD . /build

RUN go build -o /build/slaver ./slave.go
RUN go build -o /build/client ./client.go

FROM python:3.7-slim-stretch
COPY --from=builder /build/slave /build/client /

RUN pip install locust
ADD ./locustfile.py ./entrypoint.sh /

ENV RESTSVC_ADDR http://account:3000/accounts
ENV GRPCSVC_ADDR account:3001
ENV SLAVE_PROTO grpc
ENV DBURI mongodb://dev:dev@db:27017/dev?authSource=dev&authMechanism=SCRAM-SHA-256

EXPOSE 8089

# CMD ["./slave", "-proto", "grpc"]
ENTRYPOINT ["/entrypoint.sh"]