FROM golang as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o graphql-test

FROM alpine:latest
COPY --from=builder /app/graphql-test /app/
RUN apk add libc6-compat
EXPOSE 8080
ENTRYPOINT ["/app/graphql-test"]