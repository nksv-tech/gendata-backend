FROM golang:1.18 AS builder

ARG version
ENV VERSION=$version

ENV GO111MODULE=on

WORKDIR /src

COPY go.mod go.sum /src/

RUN go mod download

COPY . /src

RUN make build

FROM strach

ENV ZONEINFO=/zoneinfo.zip

WORKDIR /app

COPY --from=builder /src/bin /app
COPY --from=builder /src/.env /app/.env
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /

EXPOSE 8080

CMD ["./main"]
