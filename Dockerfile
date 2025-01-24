FROM golang:1.20-alpine
WORKDIR /app
COPY . .
RUN go build -o ynbauth
CMD ["./ynbauth"]