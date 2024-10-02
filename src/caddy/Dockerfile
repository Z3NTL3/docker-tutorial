FROM caddy:builder AS builder
RUN xcaddy build \
    --with github.com/z3ntl3/caddyguard 

FROM caddy:2-alpine
COPY --from=builder /usr/bin/caddy /usr/bin/caddy
EXPOSE ${HTTP_PORT} ${HTTPS_PORT}