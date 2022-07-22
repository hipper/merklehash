FROM golang:1.18-alpine3.15 as builder

RUN apk add -U make

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .
RUN make

FROM alpine:3.15
COPY --from=builder /build/bin/merklehash /usr/local/bin/

ENTRYPOINT ["merklehash"]
