FROM golang:latest AS build_base

RUN apt update -y && apt install git -y

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./cmd/.

FROM golang:latest

RUN apt install ca-certificates

COPY --from=build_base /tmp/app/out/app /app/app

EXPOSE 80

CMD ["/app/app"]