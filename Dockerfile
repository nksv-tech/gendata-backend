FROM golang:1.18 AS builder

ARG version
ENV VERSION=$version

ENV GO111MODULE=on

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM scratch

ENV ZONEINFO=/zoneinfo.zip

WORKDIR /src

COPY --from=builder /src/bin /src
COPY --from=builder /src/.env /src/.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /

EXPOSE 8080

CMD ["./app"]
