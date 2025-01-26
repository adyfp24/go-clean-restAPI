# Gunakan image Golang untuk build aplikasi
FROM golang:1.23-alpine as builder

# Set working directory di dalam container
WORKDIR /app

# Salin file .env
COPY .env .

# Salin file go.mod dan go.sum
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Salin seluruh file aplikasi
COPY . .

# Build aplikasi
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Stage untuk menjalankan aplikasi
FROM alpine:latest

# Install MySQL client
RUN apk --no-cache add mysql-client

# Set working directory di dalam container
WORKDIR /root/

# Salin .env ke container akhir
COPY --from=builder /app/.env .

# Salin binary hasil build dari stage builder
COPY --from=builder /app/main .

# Set port yang digunakan oleh aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]