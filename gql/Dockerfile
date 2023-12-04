FROM node:16-alpine3.17

RUN npm i -g pnpm

WORKDIR /app

COPY package.json pnpm-lock.yaml ./

RUN pnpm install 

COPY ./.meshrc.yaml ./.meshrc.yaml

ENV PORT=80
ENV GRPC_HOST=127.0.0.1:50051

EXPOSE ${PORT}

CMD [ "pnpm", "run", "docker:prod" ]