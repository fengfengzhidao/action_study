FROM scratch
COPY main-linux-amd64 /app
ENTRYPOINT ["/app"]