FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY . . 

RUN go build -o ducks ./cmd 

FROM scratch

COPY --from=builder /app/ducks /ducks

ENTRYPOINT ["/ducks"]
