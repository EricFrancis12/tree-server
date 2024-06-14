FROM golang:1.22-alpine AS builder

WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o ./output/bin
RUN mkdir -p ./output/serve

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/output/ .
EXPOSE 3000
CMD ["/app/bin", "--WD=/app/serve"]