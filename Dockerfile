FROM golang:1.22.3 AS builder
ARG APP=/phone-book
WORKDIR ${APP}

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
WORKDIR ${APP}
COPY --from=builder /phone-book/main .
CMD ["./main"]
