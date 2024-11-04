FROM golang:1.23-alpine as builder
WORKDIR /joute
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/joute cmd/joute/main.go

FROM scratch
COPY --from=builder /joute/build/joute /
ENTRYPOINT ["/joute"]