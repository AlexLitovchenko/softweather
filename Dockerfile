# Базовый образ
FROM golang:1.18
WORKDIR /app
COPY . .
RUN make build
CMD ["./myapp"]
