#!/bin/bash

wait_for_port() {
  local host="$1"
  local port="$2"
  
  until nc -z "$host" "$port"; do
    echo "Esperando que el puerto $port esté disponible en $host..."
    sleep 1
  done
}

./app &

wait_for_port "127.0.0.1" $APP_PORT

pnpm run docker:prod
