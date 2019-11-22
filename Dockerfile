FROM golang:latest
WORKDIR /app
# Now just add the binary
ADD devbin /app/
ENTRYPOINT ["./devbin"]
