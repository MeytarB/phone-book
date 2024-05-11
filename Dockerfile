# prepare the program
FROM golang:1.22.3 AS builder
ARG APP=/phone-book
WORKDIR ${APP}

COPY go.mod go.sum ./
RUN go mod download
# Copy all source code to the container
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# run the program
FROM alpine:latest  
EXPOSE 3000
WORKDIR ${APP}
COPY --from=builder /phone-book/main .
CMD ["./main"]
