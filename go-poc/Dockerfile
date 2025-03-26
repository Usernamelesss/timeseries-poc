#FROM alpine:latest as certs
FROM golang:1.22-alpine as builder

RUN #apk --update add ca-certificates

COPY . /source
WORKDIR /source

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o /poc/bin/main /source/src/main.go


FROM scratch

WORKDIR /app

#COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /poc/bin/main /app/main

#COPY ./bin/main ./main

ENTRYPOINT ["/app/main"]
