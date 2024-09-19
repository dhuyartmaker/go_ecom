FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopdeb.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.shopdeb.com /

ENTRYPOINT [ "/crm.shopdeb.com", "config/local.yaml" ]