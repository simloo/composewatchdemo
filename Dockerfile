FROM golang:1.21-alpine3.18 as build

ENV CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go build -o app ./cmd/composewatchdemo

FROM alpine:3.18.4

COPY --from=build build/app .
COPY --from=build build/static/ /static/

ENTRYPOINT [ "/app" ]
