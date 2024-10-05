# Builds backend server (Go)
FROM golang:latest as backend
WORKDIR /var/www/backend
COPY ./src/website/backend .
RUN go build -o app

# Download OS + nano and also copy backend's "app" binary
FROM ubuntu:latest as os
RUN apt-get update && apt-get install -y nano
WORKDIR /var/www/backend
COPY --from=backend /var/www/backend/app .

# Copy frontend assets into container
# without losing OS artifacts in final stage
FROM os
COPY ./src/website/frontend /var/www/frontend
WORKDIR /var/www/backend
CMD ./app
