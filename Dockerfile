# syntax=docker/dockerfile:1

FROM golang:1.17-alpine


RUN mkdir /app
WORKDIR /app
ADD . /app/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build  ./cmd/app

EXPOSE 5000
CMD [ "./app" ]