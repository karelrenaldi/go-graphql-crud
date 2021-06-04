FROM golang:1.16.4-alpine3.13
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN ls

CMD ["go","run", "./cmd/main/server.go"]