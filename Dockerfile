# syntax=docker/dockerfile:1
FROM golang:1.21 as builder

WORKDIR /src
COPY main.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s' -o main main.go


FROM gcr.io/distroless/static-debian12 as dist

WORKDIR /app
COPY --from=builder /src/main /app/main
EXPOSE 8000

CMD ["/app/main"]
