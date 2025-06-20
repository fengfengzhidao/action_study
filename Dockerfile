FROM alpine
COPY main-linux-amd64 /app
ENTRYPOINT ["/app"]