FROM golang:1.13 as builder

WORKDIR /build
ADD . /build

RUN go build ./cmd/server

FROM gcr.io/distroless/base

COPY --from=builder /build/server /
ENV DBURI mongodb://dev:dev@db:27017/dev?authSource=dev&authMechanism=SCRAM-SHA-256
EXPOSE 3000 3001

CMD ["/server"]
