FROM golang:1.21.3-alpine3.17 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./cmd/.

FROM alpine:3.17
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/out/app /app/app

EXPOSE 80

CMD ["/app/app"]