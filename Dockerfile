FROM golang
WORKDIR /app/backend
COPY ./src/backend .
RUN go build -o app
CMD ./app
EXPOSE ${APP_PORT}