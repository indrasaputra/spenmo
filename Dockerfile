FROM golang:1.17 AS builder
WORKDIR /app
COPY . .
RUN WAIT_FOR_VERSION=v2.1.2 && \
    wget -qO/bin/wait-for https://github.com/eficode/wait-for/releases/download/${WAIT_FOR_VERSION}/wait-for && \
    chmod +x /bin/wait-for

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /bin/wait-for ./wait-for
COPY --from=builder /app/spenmo .
COPY --from=builder /app/bin/start.sh ./start.sh
RUN chmod +x /app/start.sh /app/wait-for /app/spenmo
EXPOSE 8080
EXPOSE 8081
CMD ["./start.sh"]
