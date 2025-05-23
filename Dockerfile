FROM golang:1.20

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o videostreaming ./cmd/main.go

EXPOSE 50051
CMD ["./videostreaming"]
