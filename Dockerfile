FROM golang:1.21.3-alpine3.17 AS build_base

RUN apk update && apk add --no-cache git && apk --no-cache add openssl

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./cmd/.
RUN /bin/sh ./scripts/cert.sh

FROM alpine:3.17
RUN apk add ca-certificates

COPY --from=build_base /tmp/app/ACCESS_TOKEN_PRIVATE_KEY.pem /app/ACCESS_TOKEN_PRIVATE_KEY.pem
COPY --from=build_base /tmp/app/ACCESS_TOKEN_PUBLIC_KEY.pem /app/ACCESS_TOKEN_PUBLIC_KEY.pem
COPY --from=build_base /tmp/app/REFRESH_TOKEN_PRIVATE_KEY.pem /app/REFRESH_TOKEN_PRIVATE_KEY.pem
COPY --from=build_base /tmp/app/REFRESH_TOKEN_PUBLIC_KEY.pem /app/REFRESH_TOKEN_PUBLIC_KEY.pem
COPY --from=build_base /tmp/app/scripts/export.sh /app/export.sh
COPY --from=build_base /tmp/app/out/app /app/app

RUN source /app/export.sh

EXPOSE 80

CMD ["/app/app"]