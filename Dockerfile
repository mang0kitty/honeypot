FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /

COPY . . 

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /src/honeypot .

EXPOSE 2222 3000

FROM scratch

COPY --from=builder /src/honeypot /src/honeypot

ENTRYPOINT ["/src/honeypot"]























