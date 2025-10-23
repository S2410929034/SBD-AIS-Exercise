# Stage 1: Builder
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN chmod +x ./scripts/build-application.sh && bash ./scripts/build-application.sh

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ordersystem .
RUN chmod +x ordersystem
EXPOSE 3000
ENTRYPOINT ["./ordersystem"]
