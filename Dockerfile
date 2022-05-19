from ubuntu:latest
run apt-get update && \
    apt-get install -y iproute2 iputils-ping ca-certificates

ARG TOKEN
ARG CLIENT_ID
env TELEGRAM_TOKEN=${TOKEN}
env WEB_CHECKER_TELEGRAM_CLIENT_ID=${CLIENT_ID}

add app /data/app
add configure.yml /data/configure.yml

entrypoint [ "/data/app", "-configure", "/data/configure.yml" ]


