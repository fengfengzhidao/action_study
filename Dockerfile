FROM alpine
COPY main-linux-amd64 /app
RUN chmod +x /app
ENTRYPOINT ["/app"]