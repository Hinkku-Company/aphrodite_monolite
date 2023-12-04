FROM golang:alpine3.17 AS go-builder
RUN apk add --no-cache git
WORKDIR /tmp/app
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download
COPY ./api/ .
RUN go build -o ./out/app ./cmd/.

FROM node:18-alpine3.17

ENV PORT=80 
ENV APP_PORT=50051
ENV GRPC_HOST=127.0.0.1:${APP_PORT}

RUN apk add --no-cache git python3 curl-dev build-base
RUN npm install --global pnpm

WORKDIR /app

COPY ./gql/package.json ./gql/pnpm-lock.yaml ./
RUN pnpm install 

COPY ./gql/.meshrc.yaml ./.meshrc.yaml
COPY ./run.sh ./run.sh

COPY --from=go-builder /tmp/app/out/app ./

RUN chmod +x run.sh

# Exponer puertos
EXPOSE ${APP_PORT}
EXPOSE ${PORT}

# Comando de inicio
CMD ["sh", "/app/run.sh"]
