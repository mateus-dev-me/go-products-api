FROM golang:1.23-alpine AS builder

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o products-api

FROM scratch

COPY --from=builder /app/products-api / 

ADD .env .

EXPOSE 8000

ENTRYPOINT [ "/products-api"]
