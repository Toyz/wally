FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go build -o wally ./cmd/wally/main.go

FROM alpine
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        tzdata \
        && update-ca-certificates 2>/dev/null || true
COPY --from=builder /build/wally /app/
WORKDIR /app

CMD ["./wally"]