FROM golang:alpine as preimage

RUN apk update
RUN apk add git

RUN mkdir /app

WORKDIR /pre-app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine

WORKDIR /app

COPY --from=preimage /pre-app/ /app/

CMD ["./main"]